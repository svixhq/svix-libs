# This file is @generated
import typing as t
from dataclasses import dataclass

from deprecated import deprecated

from ..internal.openapi_client import models
from ..internal.openapi_client.models.integration_in import IntegrationIn
from ..internal.openapi_client.models.integration_key_out import IntegrationKeyOut
from ..internal.openapi_client.models.integration_out import IntegrationOut
from ..internal.openapi_client.models.integration_update import IntegrationUpdate
from ..internal.openapi_client.models.list_response_integration_out import (
    ListResponseIntegrationOut,
)
from .common import ApiBase, BaseOptions, serialize_params


@dataclass
class IntegrationListOptions(BaseOptions):
    # Limit the number of returned items
    limit: t.Optional[int] = None
    # The iterator returned from a prior invocation
    iterator: t.Optional[str] = None
    # The sorting order of the returned items
    order: t.Optional[models.Ordering] = None

    def _query_params(self) -> t.Dict[str, str]:
        return serialize_params(
            {
                "limit": self.limit,
                "iterator": self.iterator,
                "order": self.order,
            }
        )


@dataclass
class IntegrationCreateOptions(BaseOptions):
    idempotency_key: t.Optional[str] = None

    def _header_params(self) -> t.Dict[str, str]:
        return serialize_params(
            {
                "idempotency-key": self.idempotency_key,
            }
        )


@dataclass
class IntegrationRotateKeyOptions(BaseOptions):
    idempotency_key: t.Optional[str] = None

    def _header_params(self) -> t.Dict[str, str]:
        return serialize_params(
            {
                "idempotency-key": self.idempotency_key,
            }
        )


class IntegrationAsync(ApiBase):
    async def list(
        self, app_id: str, options: IntegrationListOptions = IntegrationListOptions()
    ) -> ListResponseIntegrationOut:
        """List the application's integrations."""
        response = await self._request_asyncio(
            method="get",
            path="/api/v1/app/{app_id}/integration",
            path_params={
                "app_id": app_id,
            },
            query_params=options._query_params(),
            header_params=options._header_params(),
        )
        return ListResponseIntegrationOut.from_dict(response.json())

    async def create(
        self,
        app_id: str,
        integration_in: IntegrationIn,
        options: IntegrationCreateOptions = IntegrationCreateOptions(),
    ) -> IntegrationOut:
        """Create an integration."""
        response = await self._request_asyncio(
            method="post",
            path="/api/v1/app/{app_id}/integration",
            path_params={
                "app_id": app_id,
            },
            query_params=options._query_params(),
            header_params=options._header_params(),
            json_body=integration_in.to_dict(),
        )
        return IntegrationOut.from_dict(response.json())

    async def get(self, app_id: str, integ_id: str) -> IntegrationOut:
        """Get an integration."""
        response = await self._request_asyncio(
            method="get",
            path="/api/v1/app/{app_id}/integration/{integ_id}",
            path_params={
                "app_id": app_id,
                "integ_id": integ_id,
            },
        )
        return IntegrationOut.from_dict(response.json())

    async def update(
        self, app_id: str, integ_id: str, integration_update: IntegrationUpdate
    ) -> IntegrationOut:
        """Update an integration."""
        response = await self._request_asyncio(
            method="put",
            path="/api/v1/app/{app_id}/integration/{integ_id}",
            path_params={
                "app_id": app_id,
                "integ_id": integ_id,
            },
            json_body=integration_update.to_dict(),
        )
        return IntegrationOut.from_dict(response.json())

    async def delete(self, app_id: str, integ_id: str) -> None:
        """Delete an integration."""
        await self._request_asyncio(
            method="delete",
            path="/api/v1/app/{app_id}/integration/{integ_id}",
            path_params={
                "app_id": app_id,
                "integ_id": integ_id,
            },
        )

    @deprecated
    async def get_key(self, app_id: str, integ_id: str) -> IntegrationKeyOut:
        """Get an integration's key."""
        response = await self._request_asyncio(
            method="get",
            path="/api/v1/app/{app_id}/integration/{integ_id}/key",
            path_params={
                "app_id": app_id,
                "integ_id": integ_id,
            },
        )
        return IntegrationKeyOut.from_dict(response.json())

    async def rotate_key(
        self,
        app_id: str,
        integ_id: str,
        options: IntegrationRotateKeyOptions = IntegrationRotateKeyOptions(),
    ) -> IntegrationKeyOut:
        """Rotate the integration's key. The previous key will be immediately revoked."""
        response = await self._request_asyncio(
            method="post",
            path="/api/v1/app/{app_id}/integration/{integ_id}/key/rotate",
            path_params={
                "app_id": app_id,
                "integ_id": integ_id,
            },
            query_params=options._query_params(),
            header_params=options._header_params(),
        )
        return IntegrationKeyOut.from_dict(response.json())


class Integration(ApiBase):
    def list(
        self, app_id: str, options: IntegrationListOptions = IntegrationListOptions()
    ) -> ListResponseIntegrationOut:
        """List the application's integrations."""
        response = self._request_sync(
            method="get",
            path="/api/v1/app/{app_id}/integration",
            path_params={
                "app_id": app_id,
            },
            query_params=options._query_params(),
            header_params=options._header_params(),
        )
        return ListResponseIntegrationOut.from_dict(response.json())

    def create(
        self,
        app_id: str,
        integration_in: IntegrationIn,
        options: IntegrationCreateOptions = IntegrationCreateOptions(),
    ) -> IntegrationOut:
        """Create an integration."""
        response = self._request_sync(
            method="post",
            path="/api/v1/app/{app_id}/integration",
            path_params={
                "app_id": app_id,
            },
            query_params=options._query_params(),
            header_params=options._header_params(),
            json_body=integration_in.to_dict(),
        )
        return IntegrationOut.from_dict(response.json())

    def get(self, app_id: str, integ_id: str) -> IntegrationOut:
        """Get an integration."""
        response = self._request_sync(
            method="get",
            path="/api/v1/app/{app_id}/integration/{integ_id}",
            path_params={
                "app_id": app_id,
                "integ_id": integ_id,
            },
        )
        return IntegrationOut.from_dict(response.json())

    def update(
        self, app_id: str, integ_id: str, integration_update: IntegrationUpdate
    ) -> IntegrationOut:
        """Update an integration."""
        response = self._request_sync(
            method="put",
            path="/api/v1/app/{app_id}/integration/{integ_id}",
            path_params={
                "app_id": app_id,
                "integ_id": integ_id,
            },
            json_body=integration_update.to_dict(),
        )
        return IntegrationOut.from_dict(response.json())

    def delete(self, app_id: str, integ_id: str) -> None:
        """Delete an integration."""
        self._request_sync(
            method="delete",
            path="/api/v1/app/{app_id}/integration/{integ_id}",
            path_params={
                "app_id": app_id,
                "integ_id": integ_id,
            },
        )

    @deprecated
    def get_key(self, app_id: str, integ_id: str) -> IntegrationKeyOut:
        """Get an integration's key."""
        response = self._request_sync(
            method="get",
            path="/api/v1/app/{app_id}/integration/{integ_id}/key",
            path_params={
                "app_id": app_id,
                "integ_id": integ_id,
            },
        )
        return IntegrationKeyOut.from_dict(response.json())

    def rotate_key(
        self,
        app_id: str,
        integ_id: str,
        options: IntegrationRotateKeyOptions = IntegrationRotateKeyOptions(),
    ) -> IntegrationKeyOut:
        """Rotate the integration's key. The previous key will be immediately revoked."""
        response = self._request_sync(
            method="post",
            path="/api/v1/app/{app_id}/integration/{integ_id}/key/rotate",
            path_params={
                "app_id": app_id,
                "integ_id": integ_id,
            },
            query_params=options._query_params(),
            header_params=options._header_params(),
        )
        return IntegrationKeyOut.from_dict(response.json())
