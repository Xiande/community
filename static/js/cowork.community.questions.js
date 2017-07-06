window.Questions = {
    serverUrl: CCConfig.L_Menu_BaseUrl + "/ajax/question",
    addQuestion: function (question) {
        $.ajax({
            async: false,
            type: "POST",
            cache: false,
            url: this.serverUrl,
            contentType: "application/json; charset=utf-8",
            data: encodeURI(JSON.stringify(question)),
            success: function (msg) {
                ccAlert(msg);
                window.location = "/";
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                ccAlert("AddQuestionError", XMLHttpRequest.statusText);
            }
        });
    },
    viewQuestion: function (questionId, callback) {
        var url = this.serverUrl + "?op=viewquestion&qid=" + questionId;
        $.ajax({
            type: "Get",
            url: url,
            cache: false,
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (data) {
                callback(data);
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                if (XMLHttpRequest.responseText == "QuestionNotExists") {
                    ccAlert("QuestionNotExists");
                    $("#main-panel").hide();
                    window.location = "/";
                }
            }
        });
    },
    getQuestionPage: function (targetID) {
        var sourceSort = GetRequestPara("sSort");
        var sourceKey = GetRequestPara("skey");
        var ssearchType = GetRequestPara("stype");
        if (ssearchType != "") {
            $(".searchType li").removeClass("active");
            $("#" + ssearchType).attr("class", "active");
        }
        $("#qkeyOne").attr("value", sourceKey);
        var langfilter = getCookie("LangFilter");
        var viewModel = function () {
            var self = this;
            self.loading = ko.observable(true);
            self.items = ko.observableArray([]);
            self.SearchType = ko.observable(ssearchType == "" ? "search" : ssearchType);
            self.KeyOne = ko.observable(sourceKey == "" ? "" : sourceKey);
            self.KeyTwo = ko.observable("");
            self.SortField = ko.observable(sourceSort == "" ? "CreateDate" : sourceSort);
            self.SelectedLang = ko.observable(langfilter == null ? "all" : langfilter);
            self.LangList = ko.observableArray([
	            { "langValue": "all", "langName": dict["FilterAll"] },
	            { "langValue": "en", "langName": dict["FilterEN"] },
	            { "langValue": "cn", "langName": dict["FilterCN"] }]);
            self.langChanged = function () {
                if (self.SelectedLang() != undefined && self.SelectedLang() != null) {
                    self.goToFirst();
                    setCookie("LangFilter", self.SelectedLang(), 'd30');
                }
            };
            paginationViewModel.apply(self, [10, function (page, pageHandler) {
                self.loading(true);
                $.ajax({
                    url: Questions.serverUrl,
                    cache: false,
                    dataType: "json",
                    data: {
                        op: "getques",
                        pIndex: page,
                        sType: self.SearchType,
                        pSize: self.pageSize,
                        keyOne: self.KeyOne,
                        keyTwo: self.KeyTwo,
                        sort: self.SortField,
                        lang: self.SelectedLang
                    },
                    success: function (data) {
                        pageHandler.call(self, data);
						if (data.result != null) {
							self.items(data.result);
						}
                        else{
							self.items([]);
						}
                        self.loading(false);
                        bindUserPresence();
                        //ProcessImn();
                    },
                    error: function (res) {

                    }
                });
            }]);
            self.Sort = function (sortField, data) {
                self.SortField = sortField;
                self.goToFirst();
                return true;
            };
            //标签页切换
            self.goToType = function (sType) {
                $(".searchType li").removeClass("active");
                $("#" + sType).attr("class", "active");
                self.SearchType = sType;
                self.KeyOne = $("#qkeyOne").attr("value");
                self.KeyTwo = $("#qkeyTwo").attr("value");
                self.goToFirst();
            };
        }
        var vm = new viewModel();
        ko.applyBindings(vm, $("#" + targetID)[0]);
        vm.goToFirst();

        $("#refreshQuestion").bind("click", function () {
            vm.KeyOne = $("#qkeyOne").attr("value");
            vm.KeyTwo = $("#qkeyTwo").attr("value");
            vm.goToFirst();
        });
        $("#searchQuesFir").bind("click", function () {
            $("#qkeyTwo").attr("value", "");
            vm.KeyOne = $("#qkeyOne").attr("value");
            vm.KeyTwo = "";
            vm.goToFirst();
        });
        $("#searchQuesSec").bind("click", function () {
            vm.KeyOne = $("#qkeyOne").attr("value");
            vm.KeyTwo = $("#qkeyTwo").attr("value");
            vm.goToFirst();
        });
        $("#qkeyTwo").bind("keypress", function () {
			if(event.keyCode == 13){
				vm.KeyOne = $("#qkeyOne").attr("value");
	            vm.KeyTwo = $("#qkeyTwo").attr("value");
	            vm.goToFirst();
			}
        });

    },
    getQuesForMyPage: function (targetID, searchType) {
        var viewModel = function () {
            var self = this;
            self.loading = ko.observable(true);
            self.items = ko.observableArray([]);
            self.SearchType = ko.observable(searchType);
            paginationViewModel.apply(self, [10, function (page, pageHandler) {
                self.loading(true);

                $.ajax({
                    url: Questions.serverUrl,
                    cache: false,
                    dataType: "json",
                    data: {
                        op: "getques",
                        pIndex: page,
                        sType: self.SearchType,
                        pSize: self.pageSize,
                        keyOne: "",
                        keyTwo: "",
                        sort: ""
                    },
                    success: function (data) {
                        pageHandler.call(self, data);
                        self.items(data.result);
                        self.loading(false);
                        bindUserPresence();
                        //ProcessImn();
                    },
                    error: function (res) {

                    }
                });
            }]);
        }
        var vm = new viewModel();
        ko.applyBindings(vm, $("#" + targetID)[0]);
        vm.goToFirst();
    },
    getFavoriteQuesPage: function (targetID) {
        var viewModel = function () {
            var self = this;
            self.loading = ko.observable(true);
            self.items = ko.observableArray([]);
            paginationViewModel.apply(self, [10, function (page, pageHandler) {
                self.loading(true);
                $.ajax({
                    url: Questions.serverUrl,
                    cache: false,
                    dataType: "json",
                    data: {
                        op: "getfavques",
                        pIndex: page,
                        pSize: self.pageSize
                    },
                    success: function (data) {
                        pageHandler.call(self, data);
                        self.items(data.result);

                        bindUserPresence();
                        //ProcessImn();
                        self.loading(false);
                    },
                    error: function (res) {

                    }
                });
            }]);
        }
        var vm = new viewModel();
        ko.applyBindings(vm, $("#" + targetID)[0]);
        vm.goToFirst();
    },
    favoriteQuestion: function (questionId, callback) {
        var url = this.serverUrl + "?op=favoritequestion&qid=" + questionId;
        $.ajax({
            type: "Get",
            url: url,
            cache: false,
            contentType: "application/json; charset=utf-8",
            success: function (data) {

                if (callback) {
                    callback(data);
                }
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                ccAlert("FavoriteQuestionError", XMLHttpRequest.statusText);
            }
        });
    },
    deleteQuestion: function (questionId, callback) {
        var url = this.serverUrl + "?op=deletequestion&qid=" + questionId;
        $.ajax({
            type: "Get",
            url: url,
            cache: false,
            contentType: "application/json; charset=utf-8",
            success: function (msg) {
                if (callback) {
                    callback(msg);
                }
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                ccAlert("DeleteQuestionError", XMLHttpRequest.statusText);
            }
        });
    },
    GetTopHotQuestion: function (targetID) {
        $.ajax({
            url: Questions.serverUrl,
            cache: false,
            dataType: "json",
            data: {
                op: "gethotques",
                top: 5
            },
            success: function (data) {
                var viewModel = function () {
                    var self = this;
                    self.items = ko.observableArray([]);
                };
                var vm = new viewModel();
                vm.items(data);
                ko.applyBindings(vm, $("#" + targetID)[0])

                bindUserPresence();
                //ProcessImn();

            },
            error: function (res) {
                debugger;
            }
        });
    },
    GetTopDynamic: function (targetID) {
        $.ajax({
            url: Questions.serverUrl,
            cache: false,
            dataType: "json",
            data: {
                op: "topdynamic",
                top: 5
            },
            success: function (data) {
                var viewModel = function () {
                    var self = this;
                    self.items = ko.observableArray([]);
                };
                var vm = new viewModel();
                vm.items(data);
                ko.applyBindings(vm, $("#" + targetID)[0])
            },
            error: function (res) {
                debugger;
            }
        });
    },
    StopfollowQuestion: function (qid, parent,data) {
        $.ajax({
            type: "get",
            url: Questions.serverUrl,
            cache: false,
            contentType: "application/x-www-form-urlencoded; charset=UTF-8",
            data: {
                op: "removefq",
                fqid: qid
            },
            success: function (result) {
                if (result == "Success") {
                    alert(dict["RemoveFavQuesSuccess"]);
                    parent.items.remove(data);
                } else {
                    alert(dict["RemoveFavQuesError"]);
                }
            },
            error: function (res) {
                alert(dict["RemoveFavQuesError"]);
                debugger;
            }
        });
    }
}
