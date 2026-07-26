GHCP := ../ghcp
OWNER := suzuki-shunsuke
REPO := ghcp
GHCP_FLAGS := --debug -u $(OWNER) -r $(REPO)

GITHUB_RUN_NUMBER ?= 0
BRANCH_PREFIX := ghcp-acceptance-test-$(GITHUB_RUN_NUMBER)
GITHUB_PR_NUMBER := $(word 3,$(subst /, ,$(GITHUB_REF)))

.PHONY: test
test:
	# create master branch
	$(GHCP) commit $(GHCP_FLAGS) -b $(BRANCH_PREFIX)-master --no-parent -m "Create master branch" fixture1.txt
	# create feature branch from master branch and add a file
	$(GHCP) commit $(GHCP_FLAGS) -b $(BRANCH_PREFIX)-feature --parent $(BRANCH_PREFIX)-master -m "Add fixture2.txt" fixture2.txt
	# add a file to feature branch
	$(GHCP) commit $(GHCP_FLAGS) -b $(BRANCH_PREFIX)-feature -m "Add fixture3.txt" --author-name="octocat" --author-email="octocat@github.com" fixture3.txt
	# add a file to feature branch but do nothing
	$(GHCP) commit $(GHCP_FLAGS) -b $(BRANCH_PREFIX)-feature -m "Add fixture1.txt" fixture1.txt
	# create a pull request from feature to master
	$(GHCP) pull-request $(GHCP_FLAGS) -b $(BRANCH_PREFIX)-feature --base $(BRANCH_PREFIX)-master \
		--title "ghcp: acceptance-test $(BRANCH_PREFIX)" \
		--body "ref: #$(GITHUB_PR_NUMBER)"

# The workflow checks out with persist-credentials: false, so git has no
# credentials to push with. Use the API with GITHUB_TOKEN instead.
# Closing the pull request deletes the feature branch; the master branch is not
# a pull request head, so it has to be deleted on its own.
.PHONY: clean-up
clean-up:
	-gh pr close $(BRANCH_PREFIX)-feature --repo $(OWNER)/$(REPO) --delete-branch
	-gh api -X DELETE repos/$(OWNER)/$(REPO)/git/refs/heads/$(BRANCH_PREFIX)-master
	-gh api -X DELETE repos/$(OWNER)/$(REPO)/git/refs/heads/$(BRANCH_PREFIX)-feature
