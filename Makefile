# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
GO_VERSION=1.12.17
PRJDIR=../paligo/
ifdef GITLAB_CI
	# set environment variables on local machine or GitLab CI
	export PRJDIR=$(CURDIR)
	export GOROOT=$(realpath $(PRJDIR)/go)
	export GOPATH=$(realpath $(PRJDIR))
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif
ifdef GITHUB_ACTIONS
	export GOPATH=$(realpath $(CURDIR))
endif

PALILIB=$(GOPATH)/src/github.com/siongui/gopalilib/lib
PALIUTIL=$(GOPATH)/src/github.com/siongui/gopalilib/util
DATA_REPO_DIR=$(CURDIR)/data
VFSDIR=$(GOPATH)/src/pali/words/vfspkg
LOCALE_DIR=$(CURDIR)/locale
TIPITAKA_XML_DIR=/tmp/tpkxml/


current_working_target: test_lib_tipitaka

###############################
# Common library for frontend #
###############################
test_libfrontend: fmt
	@echo "\033[92mTesting common library for frontend ...\033[0m"
	@cd libfrontend; echo "FIXME: do not know how to test"
###############################
# Common library for frontend #
###############################


##########################################################
# Common library for online/offline, dictionary/tipitaka #
##########################################################
test_lib: test_filter test_string test_lib_url_dictionary test_lib_tipitaka test_lib_trie test_lib_dicmgr test_lib_gettext test_lib_jsgettext
	@echo "\033[92mTesting common library for online/offline dictionary/tipitaka ...\033[0m"
	@cd lib; go test -v $(shell cd lib; ls *.go)

test_filter: fmt
	@echo "\033[92mTesting filter methods in common library for online/offline dictionary/tipitaka ...\033[0m"
	@cd lib; go test -v dictionary.go filter.go setting.go filter_test.go

test_string: fmt
	@echo "\033[92mTesting string methods in common library for online/offline dictionary/tipitaka ...\033[0m"
	@cd lib; go test -v string.go string_test.go

test_lib_tipitaka: fmt
	@echo "\033[92mTesting tipitaka methods in common library for online/offline dictionary/tipitaka ...\033[0m"
	@cd lib/tipitaka; go test -v $(shell cd lib/tipitaka; ls *.go)

test_lib_trie: fmt
	@echo "\033[92mTesting trie methods in common library for online/offline dictionary/tipitaka ...\033[0m"
	@cd lib/trie; go test -v trie.go trie_test.go
	@cd lib/trie; go test -v trie.go savetrie_test.go
	@cd lib/trie; go test -v trie.go loadtrie_test.go

test_lib_url_dictionary: fmt
	@echo "\033[92mTesting url methods in common library for online/offline dictionary/tipitaka ...\033[0m"
	@cd lib/dictionary; go test -v url.go url_test.go

test_lib_dicmgr: fmt
	@echo "\033[92mTesting dictionary manager in common library for online/offline dictionary/tipitaka ...\033[0m"
	@cd lib/dicmgr/; go test -v

test_lib_gettext: fmt
	@echo "\033[92mTesting gettext in common library for online/offline dictionary/tipitaka ...\033[0m"
	@cd lib/gettext/; go test -v -args -localeDir=$(LOCALE_DIR)

test_lib_jsgettext: fmt
	@echo "\033[92mTesting jsgettext in common library for online/offline dictionary/tipitaka ...\033[0m"
	@cd lib/jsgettext/; go test -v
#################################################################
# End of Common library for online/offline, dictionary/tipitaka #
#################################################################


##############################
# Bootstrap/Setup Dictionary #
##############################
# input file path for testing
BookCSV=$(DATA_REPO_DIR)/dictionary/dict-books.csv
WordCSV1=$(DATA_REPO_DIR)/dictionary/dict_words_1.csv
WordCSV2=$(DATA_REPO_DIR)/dictionary/dict_words_2.csv
# output file path for testing
OUTPUT_DIR=/tmp/pali/
OUTPUT_METADATA_DIR=$(OUTPUT_DIR)/metadata/
OutputBookJSON=$(OUTPUT_METADATA_DIR)/BookIdAndInfos.json
OUTPUT_PALI_WORDS_JSON_DIR=$(OUTPUT_DIR)/json/
TrieJSON=$(OUTPUT_METADATA_DIR)/trie.json
OUTPUT_METADATA_GO_FILE=$(CURDIR)/lib/dicmgr/data.go
OUTPUT_POJSON_GO_FILE=$(CURDIR)/lib/jsgettext/data.go

test_dictionary: test_po2json test_bookparser test_wordparser test_triebuild test_vfsbuild test_symlink test_embedmetadata test_check_compile

test_po2json: fmt
	@echo "\033[92mTesting converting PO files to JSON file ...\033[0m"
	cd dicutil; go test -v po2json.go po2json_test.go -args -localeDir=$(LOCALE_DIR) -outputGoDataFilePath=$(OUTPUT_POJSON_GO_FILE)
	@make fmt

test_bookparser: fmt
	@echo "\033[92mTesting parse CSV of dictionary books ...\033[0m"
	@cd dicutil; go test -v bookparser.go bookparser_test.go -args -BookCSV=$(BookCSV) -OutputBookJSON=$(OutputBookJSON)

test_wordparser: fmt
	@echo "\033[92mTesting parse CSV of dictionary words ...\033[0m"
	#@[ -d /tmp/paliwords/ ] || mkdir /tmp/paliwords/
	@cd dicutil; go test -v wordparser.go wordparser_test.go -args -WordCSV1=$(WordCSV1) -WordCSV2=$(WordCSV2) -wordsJsonDir=$(OUTPUT_PALI_WORDS_JSON_DIR)

# test_triebuild must run after test_wordparser
test_triebuild: fmt
	@echo "\033[92mTesting building succinct trie ...\033[0m"
	@cd dicutil; go test -v triebuild.go triebuild_test.go -args -wordsJsonDir=$(OUTPUT_PALI_WORDS_JSON_DIR) -trieJson=$(TrieJSON)

# test_vfsbuild must run after test_wordparser
test_vfsbuild: fmt
	@echo "\033[92mBuilding virtual file system of Pāli dictionary words ...\033[0m"
	#@[ -d $(VFSDIR) ] || mkdir -p $(VFSDIR)
	@cd dicutil; go test -v vfsbuild.go vfsbuild_test.go -args -pkgdir=$(VFSDIR) -wordsJsonDir=$(OUTPUT_PALI_WORDS_JSON_DIR)
	@cd dicutil; go test -v vfs_test.go -args -wordsJsonDir=$(OUTPUT_PALI_WORDS_JSON_DIR)

# test_symlink must run after test_vfsbuild
test_symlink: fmt
	@echo "\033[92mTesting making Pāli Dictionary symlinks for GitHub Pages...\033[0m"
	@cd dicutil; go test -v symlink.go symlink_test.go -args -outputDir=$(OUTPUT_DIR)

# run after test_bookparser and test_triebuild
test_embedmetadata: fmt
	@echo "\033[92mTesting embed metadata into Go code...\033[0m"
	@cd dicutil; go test -v embedmetadata.go embedmetadata_test.go -args -metadataDir=$(OUTPUT_METADATA_DIR) -outputGoFilePath=$(OUTPUT_METADATA_GO_FILE)
	@make fmt

test_check_compile: fmt
	@echo "\033[92mTesting checking compile of dicutil...\033[0m"
	cd dicutil; go test -v $(shell cd dicutil; ls | grep -v _test.go)

test_extract_one_dic: fmt
	@echo "\033[92mTesting extracting one dictionary...\033[0m"
	cd dicutil; go test -v onedic.go onedic_test.go -args -WordCSV1=$(WordCSV1) -WordCSV2=$(WordCSV2)
#####################################
# End of Bootstrap/Setup Dictionary #
#####################################


############################
# Bootstrap/Setup Tipiṭaka #
############################
test_tipitaka: test_build_tpk_tree

test_download_tpk: fmt
	@echo "\033[92mTesting download Tipiṭaka xml from https://www.tipitaka.org/romn/ ...\033[0m"
	@cd tpkutil; go test -v downloadtpk.go downloadtpk_test.go

clone_tpk_xml:
	@echo "\033[92mClone Tipiṭaka XML repo ...\033[0m"
	@[ -d $(TIPITAKA_XML_DIR) ] || git clone https://github.com/siongui/tipitaka-romn.git $(TIPITAKA_XML_DIR)

test_build_tpk_tree: fmt clone_tpk_xml
	@echo "\033[92mTesting build Tipiṭaka tree ...\033[0m"
	@cd tpkutil; go test -v buildtpktree.go buildtpktree_test.go -args -tpkXmlDir=$(TIPITAKA_XML_DIR)
###################################
# End of Bootstrap/Setup Tipiṭaka #
###################################


##############
# Misc Tools #
##############
test_util: fmt
	@echo "\033[92mTest utility func ...\033[0m"
	@cd util; go test -v
#####################
# End of Misc Tools #
#####################


###################
# Install Library #
###################
install: install_palilib install_goef

install_palilib:
	#go get -u github.com/siongui/gopalilib/lib
	#go get -u github.com/siongui/gopalilib/lib/trie
	go get -u github.com/siongui/gopalilib/lib/dictionary
	go get -u github.com/siongui/gopalilib/lib/dicmgr
	go get -u github.com/siongui/gopalilib/util

install_goef:
	@echo "\033[92mInstalling Go file embedder ...\033[0m"
	go get -u github.com/siongui/goef

install_local:
	@echo "\033[92mInstall ${PALILIB} locally ...\033[0m"
	@rm -rf ${PALILIB}
	@mkdir -p ${PALILIB}
	@cp -r lib/* ${PALILIB}/
	@echo "\033[92mInstall ${PALIUTIL} locally ...\033[0m"
	@rm -rf ${PALIUTIL}
	@mkdir -p ${PALIUTIL}
	@cp -r util/* ${PALIUTIL}/

# installed by go get -u github.com/siongui/gopalilib/lib (install_palilib)
lib_succinct_trie:
	@echo "\033[92mInstalling Go Succinct Trie library ...\033[0m"
	go get -u github.com/siongui/go-succinct-data-structure-trie

# installed by go get -u github.com/siongui/gopalilib/util (install_palilib)
install_gojianfan:
	@echo "\033[92mInstalling Go Chinese conversion package ...\033[0m"
	go get -u github.com/siongui/gojianfan

# installed by go get -u github.com/siongui/gopalilib/util (install_palilib)
install_charset:
	@echo "\033[92mInstalling golang.org/x/net/html/charset ...\033[0m"
	go get -u golang.org/x/net/html/charset

# installed by go get -u github.com/siongui/gopalilib/util (install_palilib)
install_rabbit:
	@echo "\033[92mInstalling github.com/winhtaikaung/Rabbit-Go/rabbit ...\033[0m"
	go get -u github.com/winhtaikaung/Rabbit-Go/rabbit
##########################
# End of Install Library #
##########################


########
# Misc #
########
fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt lib/*.go
	@go fmt lib/dictionary/*.go
	@go fmt lib/tipitaka/*.go
	@go fmt lib/trie/*.go
	@go fmt lib/dicmgr/*.go
	@go fmt lib/gettext/*.go
	@go fmt lib/jsgettext/*.go
	@go fmt dicutil/*.go
	@go fmt tpkutil/*.go
	@go fmt util/*.go
	@go fmt libfrontend/*.go
	@go fmt libfrontend/everyword/*.go
	@go fmt libfrontend/treeview/*.go
	@go fmt libfrontend/xslt/*.go

clean:
	rm -rf pkg/ src/

clone_pali_data:
	@echo "\033[92mClone Pāli data Repo ...\033[0m"
	@[ -d $(DATA_REPO_DIR) ] || git clone https://github.com/siongui/data.git $(DATA_REPO_DIR) --depth=1

download_go:
	@echo "\033[92mDownloading and Installing Go ...\033[0m"
	@wget https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz
	@tar -xvzf go$(GO_VERSION).linux-amd64.tar.gz
	@rm go$(GO_VERSION).linux-amd64.tar.gz
