# TODO - remove this special case when we fix the generated code for empty openapi structs
from typing import Any, Dict, TypeVar

T = TypeVar("T", bound="EventTypeFromOpenApiSchemas")

EventTypeFromOpenApiSchemas = Dict[str, Any]
