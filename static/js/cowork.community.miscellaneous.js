
window.Miscellaneous = {
    serverUrl: CCConfig.L_Menu_BaseUrl + "/ajax/common",
    getFixedCount: function (callback) {
        $.ajax({
            async: false,
            type: "get",
            cache: false,
            url: this.serverUrl + "?op=getfixedcount",
            contentType: "application/json; charset=utf-8",
            success: function (count) {

                if (!isNaN(count) && callback) {
                    callback(count);
                }
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                debugger;
            }
        });
    },
    GetPopularSite: function (targetID) {
        var viewModel = function () {
            var self = this;
            self.loading = ko.observable(true);
            self.items = ko.observableArray([]);
            paginationViewModel.apply(self, [10, function (page, pageHandler) {
                $.ajax({
                    url: Miscellaneous.serverUrl,
                    cache: false,
                    dataType: "json",
                    data: {
                        op: "popularsite",
                        topcount: 20,
                        radom: 5
                    },
                    success: function (data) {
                        pageHandler.call(self, data);
                        self.items(data);
                        self.loading(false);
                    },
                    error: function (res) {
                        debugger;
                    }
                });
            }]);
        }
        var popSite = new viewModel();
        ko.applyBindings(popSite, $("#" + targetID)[0]);
        popSite.goToFirst();

        $("#refpopularSite").bind("click", function () {
            popSite.goToFirst();
        });
    },
    getStarUser: function (targetID) {
        $.ajax({
            url: Miscellaneous.serverUrl,
            cache: false,
            dataType: "json",
            data: {
                op: "staruser"
            },
            success: function (data) {
                var vm = ko.mapping.fromJS(data);
                ko.applyBindings(vm, $("#" + targetID)[0]);
            },
            error: function (res) {
                debugger;
            }
        });
    },
    StopFollowing: function (caseType, id, parent, data) {
        switch (caseType) {
            case "question":
                Questions.StopfollowQuestion(id, parent, data);
                break;
            case "answer":
                Answers.StopfollowAnswer(id, parent, data);
                break;
            default:
                alert(dict["SystemError"]);
                break;
        }
    },
    AddGlossary: function () {
        var glossary = $("#glossary").val();
        var fname = $("#fname").val();
        if (glossary == "") {
            alert(dict["GlossaryMust"]);
            return;
        }
        if (fname == "") {
            alert(dict["FullNameMust"]);
            return;
        }
        var gloadata =  {
                op: "addglossary",
                glossary: glossary,
                fullname: fname,
                desc: $("#desc").val(),
                scope: $("#scope").val(),
                localname: $("#lname").val()
            };
        $.ajax({
            url: Miscellaneous.serverUrl + "?op=addglossary",
            async: false,
            type: "post",
            cache: false,
            contentType: "application/json; charset=utf-8",
            data : encodeURI(JSON.stringify(gloadata)),
            success: function (data) {
                //alert(data);
                if (data == "Success") {
                    $("#askgdiv").hide();
                    $("#finishdiv").show();
                    $("#searchDict").attr("href", "http://my.lenovo.com/my/DictSearch.aspx?k=" + glossary)
                } else {
                    alert(data);
                }
            },
            error: function (res) {
                debugger;
            }
        });
    },
    AddTransation: function () {
        var trantitle = $("#Transation").val();
        if (trantitle == "") {
            ccAlert("TransationMust", "");
            return;
        }

        var helpLink = $("#Helplink").val();
        if (!IsURL(helpLink)) {
            ccAlert("IncorrectHelpLink", "");
            return;
        }
        var sytemLink = $("#Systemlink").val();
        if (!IsURL(sytemLink)) {
            ccAlert("IncorrectSysLink", "");
            return;
        }
        var valid = true;
        //if($("#Validity").attr("checked") == true || $("#Validity").attr("checked") == "checked"){
        //valid = true;
        //}
        var trandata = {
                op: "addtransation",
                trantitle: trantitle,
                desc: $("#Description").val(),
                Scope: $("#Scope").val(),
                Hlink: helpLink,
                Slink: sytemLink,
                Depart: $("#Department").val(),
                info: $("#Contactinformation").val(),
                keywords: $("#Keywords").val(),
                Addmarks: $("#Additionalremarks").val(),
                Validity: valid
            };
        $.ajax({
            url: Miscellaneous.serverUrl + "?op=addtransation",
            async: false,
            type: "post",
            cache: false,
            contentType: "application/json; charset=utf-8",
            data: encodeURI(JSON.stringify(trandata)),
            success: function (data) {
                //alert(data);
                if (data == "Success") {
                    $("#askgdiv").hide();
                    $("#finishdiv").show();
                    $("#searchDict").attr("href", "http://my.lenovo.com/my/DrSearch.aspx?k=" + trantitle)
                } else {
                    alert(data);
                }
            },
            error: function (res) {
                debugger;
            }
        });
    },
    Position: function (index) {
        return Miscellaneous.PostionTemp(index);
    },
    PostionTemp: function (index) {
        var positionList = new Array();
        positionList.push([{ Top: 60, Left: 280 }, { Top: 80, Left: 720 }, { Top: 220, Left: 110 }, { Top: 240, Left: 560 }, { Top: 360, Left: 250 }]);
        positionList.push([{ Top: "10%", Left: 280 }, { Top: "30%", Left: 560 }, { Top: "40%", Left: 120 }, { Top: "55%", Left: 760 }, { Top: "60%", Left: 320 }]);
        positionList.push([{ Top: 30, Left: 280 }, { Top: 70, Left: 760 }, { Top: 170, Left: 120 }, { Top: 240, Left: 760 }, { Top: 300, Left: 320 }]);
        positionList.push([{ Top: 30, Left: 280 }, { Top: 70, Left: 760 }, { Top: 170, Left: 120 }, { Top: 240, Left: 760 }, { Top: 300, Left: 320 }]);
        positionList.push([{ Top: 30, Left: 280 }, { Top: 70, Left: 760 }, { Top: 170, Left: 120 }, { Top: 240, Left: 760 }, { Top: 300, Left: 320 }]);

        return positionList[index];
    },
    RandomHeight: function () {
        return Math.floor(Math.random()*(140-60+1)+60);
    },
    ThankWall: function (targetID) {
    	var lang = "English";
    	if(getCookie("CurLang") == "zh-cn"){
    		lang = "Chinese";
    	}
        var viewModel = function () {
            var self = this;
            self.loading = ko.observable(true);
            self.items = ko.observableArray([]);
            self.PTIndex = ko.observable(1);
            self.PagingInfo = ko.observable("");
            self.PrePaging = ko.observable("");
            self.NextPaging = ko.observable("");
            paginationViewModel.apply(self, [5, function (page, pageHandler) {
                self.loading(true);
                $.ajax({
                    url: Miscellaneous.serverUrl,
                    type: "get",
                    cache: false,
                    dataType: "json",
                    contentType: "application/json; charset=utf-8",
                    data: {
                        op: "thkwall",
                        pSize: self.pageSize,
                        pIndex: page,
                        lang : lang,
                        nextpage : self.PagingInfo()
                    },
                    success: function (data) {
	                    pageHandler.call(self, data);
	                    self.items(data.data);
	                    self.PrePaging(data.previnfo);
            			self.NextPaging(data.nextinfo)
	                    //self.PTIndex(Math.ceil(Math.random() * 5) - 1);
                    },
                    error: function (res) {
                        debugger;
                    }
                });
            }]);
		    self.goPre = function () {
		        var cur = self.currentPage();
		        if (cur > 1) {
		        	self.PagingInfo(self.PrePaging());
		        	self.goToPage(cur - 1);
		        };
		    };
		    self.goNext = function () {
		        var cur = self.currentPage();
		        if (cur < self.pageCount()) {
		        	self.PagingInfo(self.NextPaging());
		        	self.goToPage(cur + 1);
		        };
		    };
			
			self.showNextElement = function(elem) { 
				$(elem).hide().slideDown(1000); 
			};
    		self.hidePreElement = function(elem) { 
    			$(elem).slideUp(1000,function() { 
    				$(elem).remove(); 
    			});
    		}
        }
        var vm = new viewModel();
        ko.applyBindings(vm, $("#" + targetID)[0]);
        vm.goToFirst();
    }

}

