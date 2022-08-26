build: git-submodules
	@echo "Building..."
	@rm -rf add.sh del.sh china_ip_list.txt
	cp china_ip_list/china_ip_list.txt ./
	go build .
	./wg-route

git-submodules:
	@[ -d ".git" ] || (echo "Not a git repository" && exit 1)
	@echo "Updating git submodules"
	@# Dockerhub using ./hooks/post-checkout to set submodules, so this line will fail on Dockerhub
	@# these lines will also fail if ran as root in a non-root user's checked out repository
	@git submodule sync --quiet --recursive || true
	@git submodule update --quiet --init --recursive --force || true