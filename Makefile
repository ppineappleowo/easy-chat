user-rpc-dev:
	@make -f deploy/mk/user-rpc.mk release-teat

user-api-dev:
	@make -f deploy/mk/user-api.mk release-teat

release-test:user-rpc-dev user-api-dev

install-server:
	cd ./deploy/script && chmod +x release-test.sh && ./release-test.sh