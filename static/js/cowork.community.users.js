window.Users = {
    serverUrl: CCConfig.L_Menu_BaseUrl + "/ajax/user",
    followUser: function (itcode,IsFollowed, callback) {
    	var operation = "followuser";
    	if(IsFollowed){
    		operation = "stopfollowuser";
    	}
        $.ajax({
            url: Users.serverUrl,
            cache: false,
            data: {
                op: operation,
                itcode: itcode
            },
            success: function (msg) {
                
                if (callback != undefined) {
                    callback(msg);
                }
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                ccAlert("FollowUserError", XMLHttpRequest.statusText);
            }
        });
    },
    GetStatisticsInfo: function (targetID) {
        $.ajax({
            url: Users.serverUrl,
            cache: false,
            dataType: "json",
            data: {
                op: "statistics"
            },
            success: function (data) {
                var viewModel = function () {
                    FavoriteCount = data.FavoriteCount;
                    BestCount = data.BestCount;
                    TagsCount = data.TagsCount;
                    ViewCount = data.ViewCount;
                    VotedCount = data.VotedCount;
                    PhotoStr = data.PhotoStr;
                    UserName = data.UserName;
                    BadgesCount = data.BadgesCount;
                    ExpertCount = data.ExpertCount;
                };
                var vm = new viewModel();
                ko.applyBindings(vm, $("#" + targetID)[0])
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
				debugger
            }
        });
    },
    GetPoints: function (targetID) {
        $.ajax({
            url: Users.serverUrl,
            cache: false,
            data: {
                op: "getpoints"
            },
            success: function (data) {
                $("#" + targetID).text(data.Currency());
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {

            }
        });
    }

}

