// SPDX-FileCopyrightText: © 2022 Svix Authors
// SPDX-License-Identifier: MIT

use reqwest::StatusCode;

use svix_server::{
    core::types::MessageStatus,
    v1::{
        endpoints::attempt::{AttemptedMessageOut, MessageAttemptOut},
        utils::ListResponse,
    },
};

mod utils;

use utils::{
    common_calls::{
        create_test_app, create_test_endpoint, create_test_message,
        get_msg_attempt_list_and_assert_count,
    },
    get_default_test_config, run_with_retries, start_svix_server, start_svix_server_with_cfg,
    TestReceiver,
};

use std::time::Duration;

#[tokio::test]
async fn test_list_attempted_messages() {
    let (client, _jh) = start_svix_server();

    let app_id = create_test_app(&client, "app1").await.unwrap().id;

    let receiver_1 = TestReceiver::start(axum::http::StatusCode::OK);
    let receiver_2 = TestReceiver::start(axum::http::StatusCode::OK);

    let endp_id_1 = create_test_endpoint(&client, &app_id, &receiver_1.endpoint)
        .await
        .unwrap()
        .id;
    let endp_id_2 = create_test_endpoint(&client, &app_id, &receiver_2.endpoint)
        .await
        .unwrap()
        .id;

    let msg_1 = create_test_message(&client, &app_id, serde_json::json!({"test": "data1"}))
        .await
        .unwrap();
    let msg_2 = create_test_message(&client, &app_id, serde_json::json!({"test": "data2"}))
        .await
        .unwrap();
    let msg_3 = create_test_message(&client, &app_id, serde_json::json!({"test": "data3"}))
        .await
        .unwrap();

    let list_1: ListResponse<AttemptedMessageOut> = client
        .get(
            &format!("api/v1/app/{}/endpoint/{}/msg/", app_id, endp_id_1),
            StatusCode::OK,
        )
        .await
        .unwrap();
    let list_2: ListResponse<AttemptedMessageOut> = client
        .get(
            &format!("api/v1/app/{}/endpoint/{}/msg/", app_id, endp_id_2),
            StatusCode::OK,
        )
        .await
        .unwrap();

    for list in [list_1, list_2] {
        assert_eq!(list.data.len(), 3);

        // Assert order
        assert_eq!(list.data[0].msg, msg_3);
        assert_eq!(list.data[1].msg, msg_2);
        assert_eq!(list.data[2].msg, msg_1);
    }
}

#[tokio::test]
async fn test_list_attempts_by_endpoint() {
    let (client, _jh) = start_svix_server();

    let app_id = create_test_app(&client, "v1AttemptListAttemptsByEndpointTestApp")
        .await
        .unwrap()
        .id;

    let receiver_1 = TestReceiver::start(axum::http::StatusCode::OK);
    let receiver_2 = TestReceiver::start(axum::http::StatusCode::OK);

    let endp_id_1 = create_test_endpoint(&client, &app_id, &receiver_1.endpoint)
        .await
        .unwrap()
        .id;
    let endp_id_2 = create_test_endpoint(&client, &app_id, &receiver_2.endpoint)
        .await
        .unwrap()
        .id;

    let msg_1 = create_test_message(&client, &app_id, serde_json::json!({"test": "data1"}))
        .await
        .unwrap();
    let msg_2 = create_test_message(&client, &app_id, serde_json::json!({"test": "data2"}))
        .await
        .unwrap();
    let msg_3 = create_test_message(&client, &app_id, serde_json::json!({"test": "data3"}))
        .await
        .unwrap();

    // And wait at most one second for all attempts to be processed
    run_with_retries(|| async {
        for endp_id in [endp_id_1.clone(), endp_id_2.clone()] {
            let list: ListResponse<MessageAttemptOut> = client
                .get(
                    &format!("api/v1/app/{}/attempt/endpoint/{}/", app_id, endp_id),
                    StatusCode::OK,
                )
                .await
                .unwrap();

            if list.data.len() != 3 {
                anyhow::bail!("list len {}, not 3", list.data.len());
            }
        }

        Ok(())
    })
    .await
    .unwrap();

    let list_1: ListResponse<MessageAttemptOut> = client
        .get(
            &format!("api/v1/app/{}/attempt/endpoint/{}/", app_id, endp_id_1),
            StatusCode::OK,
        )
        .await
        .unwrap();
    let list_2: ListResponse<MessageAttemptOut> = client
        .get(
            &format!("api/v1/app/{}/attempt/endpoint/{}/", app_id, endp_id_2),
            StatusCode::OK,
        )
        .await
        .unwrap();

    for list in [list_1, list_2] {
        let message_ids: Vec<_> = list.data.into_iter().map(|amo| amo.msg_id).collect();
        assert!(message_ids.contains(&msg_1.id));
        assert!(message_ids.contains(&msg_2.id));
        assert!(message_ids.contains(&msg_3.id));
    }

    receiver_1.jh.abort();
    receiver_2.jh.abort();
}

#[tokio::test]
async fn test_message_attempts() {
    let mut cfg = get_default_test_config();
    cfg.retry_schedule = (0..2)
        .into_iter()
        .map(|_| Duration::from_millis(1))
        .collect();

    let (client, _jh) = start_svix_server_with_cfg(&cfg);

    for (status_code, msg_status, attempt_count) in [
        // Success
        (StatusCode::OK, MessageStatus::Success, Some(1)),
        // HTTP 400
        (StatusCode::FORBIDDEN, MessageStatus::Fail, None),
        // HTTP 500
        (StatusCode::INTERNAL_SERVER_ERROR, MessageStatus::Fail, None),
    ] {
        let app_id = create_test_app(&client, "app").await.unwrap().id;

        let receiver = TestReceiver::start(status_code);

        let endp_id = create_test_endpoint(&client, &app_id, &receiver.endpoint)
            .await
            .unwrap()
            .id;

        let msg = create_test_message(&client, &app_id, serde_json::json!({"test": "data"}))
            .await
            .unwrap();

        let list = get_msg_attempt_list_and_assert_count(
            &client,
            &app_id,
            &msg.id,
            attempt_count.unwrap_or(&cfg.retry_schedule.len() + 1),
        )
        .await
        .unwrap();

        for i in list.data.iter() {
            assert_eq!(i.status, msg_status);
            println!("{} {}", i.response_status_code, status_code);
            assert_eq!(
                i.response_status_code,
                TryInto::<i16>::try_into(status_code.as_u16()).unwrap()
            );
            assert_eq!(i.endpoint_id, endp_id);
        }
        receiver.jh.abort();
    }

    // non-HTTP-related failures:
    let app_id = create_test_app(&client, "app").await.unwrap().id;

    let receiver = TestReceiver::start(StatusCode::OK);

    // stop receiver before beginning tests:
    receiver.jh.abort();

    let endp_id = create_test_endpoint(&client, &app_id, &receiver.endpoint)
        .await
        .unwrap()
        .id;

    let msg = create_test_message(&client, &app_id, serde_json::json!({"test": "data1"}))
        .await
        .unwrap();

    let list = get_msg_attempt_list_and_assert_count(
        &client,
        &app_id,
        &msg.id,
        &cfg.retry_schedule.len() + 1,
    )
    .await
    .unwrap();

    for i in list.data.iter() {
        assert_eq!(i.status, MessageStatus::Fail);
        assert_eq!(i.response_status_code, 0);
        assert_eq!(i.endpoint_id, endp_id);
    }
}
