package com.svix.kotlin

import com.svix.kotlin.exceptions.ApiException
import com.svix.kotlin.internal.apis.MessageApi
import com.svix.kotlin.models.ListResponseMessageOut
import com.svix.kotlin.models.MessageIn
import com.svix.kotlin.models.MessageOut

class Message internal constructor(debugUrl: String) {
    val api = MessageApi(debugUrl)

    suspend fun list(appId: String, options: MessageListOptions): ListResponseMessageOut {
        try {
            return api.listMessagesApiV1AppAppIdMsgGet(appId, options.iterator, options.limit, options.eventTypes, options.before)
        } catch (e: Exception) {
            throw ApiException.wrapInternalApiException(e)
        }
    }

    suspend fun create(appId: String, messageIn: MessageIn): MessageOut {
        try {
            return api.createMessageApiV1AppAppIdMsgPost(appId, messageIn)
        } catch (e: Exception) {
            throw ApiException.wrapInternalApiException(e)
        }
    }

    suspend fun get(msgId: String, appId: String): MessageOut {
        try {
            return api.getMessageApiV1AppAppIdMsgMsgIdGet(msgId, appId)
        } catch (e: Exception) {
            throw ApiException.wrapInternalApiException(e)
        }
    }
}
