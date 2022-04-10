#!/bin/sh

cd "$(dirname "$0")/.."

echo "Mock Generation has been started"

mockgen -source=repository/repository.go -destination=repository/mocks/repository_mock.go -package=repositorymocks
mockgen -source=service/service.go -destination=service/mocks/service_mock.go -package=servicemocks
mockgen -source=cache/cache.go -destination=cache/mocks/cache_mock.go -package=cachemocks

echo "Mock Generation has been finished"