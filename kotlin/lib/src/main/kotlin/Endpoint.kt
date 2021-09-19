package com.svix.kotlin

import com.svix.kotlin.exceptions.ApiException
import com.svix.kotlin.internal.apis.EndpointApi
import com.svix.kotlin.internal.infrastructure.ClientException
import com.svix.kotlin.internal.infrastructure.ServerException
import com.svix.kotlin.models.EndpointIn
import com.svix.kotlin.models.EndpointOut
import com.svix.kotlin.models.ListResponseEndpointOut

class EndpointListOptions() : ListOptions() {
    override fun iterator(iterator: String) = apply { super.iterator(iterator) }

    override fun limit(limit: Int) = apply { super.limit(limit) }
}

class Endpoint internal constructor(debugUrl: String) {
    val api = EndpointApi(debugUrl)

    suspend fun list(appId: String, options: EndpointListOptions): ListResponseEndpointOut {
        try {
            return api.listEndpointsApiV1AppAppIdEndpointGet(appId, options.iterator, options.limit)
        } catch (e: Exception) {
            when (e) {
                is ServerException, is ClientException, is UnsupportedOperationException -> {
                    throw ApiException(e)
                }
                else -> throw e
            }
        }
    }

    suspend fun create(appId: String, applicationIn: EndpointIn): EndpointOut {
        try {
            return api.createEndpointApiV1AppAppIdEndpointPost(appId, applicationIn)
        } catch (e: Exception) {
            when (e) {
                is ServerException, is ClientException, is UnsupportedOperationException -> {
                    throw ApiException(e)
                }
                else -> throw e
            }
        }
    }

    suspend fun get(appId: String, endpointIn: String): EndpointOut {
        try {
            return api.getEndpointApiV1AppAppIdEndpointEndpointIdGet(appId, endpointIn)
        } catch (e: Exception) {
            when (e) {
                is ServerException, is ClientException, is UnsupportedOperationException -> {
                    throw ApiException(e)
                }
                else -> throw e
            }
        }
    }

    suspend fun update(appId: String, endpointId: String, endpointIn: EndpointIn): EndpointOut {
        try {
            return api.updateEndpointApiV1AppAppIdEndpointEndpointIdPut(appId, endpointId, endpointIn)
        } catch (e: Exception) {
            when (e) {
                is ServerException, is ClientException, is UnsupportedOperationException -> {
                    throw ApiException(e)
                }
                else -> throw e
            }
        }
    }

    suspend fun delete(appId: String, endpointId: String) {
        try {
            api.deleteEndpointApiV1AppAppIdEndpointEndpointIdDelete(appId, endpointId)
        } catch (e: Exception) {
            when (e) {
                is ServerException, is ClientException, is UnsupportedOperationException -> {
                    throw ApiException(e)
                }
                else -> throw e
            }
        }
    }

    suspend fun getSecret(appId: String, endpointId: String) {
        try {
            api.getEndpointSecretApiV1AppAppIdEndpointEndpointIdSecretGet(appId, endpointId)
        } catch (e: Exception) {
            when (e) {
                is ServerException, is ClientException, is UnsupportedOperationException -> {
                    throw ApiException(e)
                }
                else -> throw e
            }
        }
    }
}
