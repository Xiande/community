/*
URL和Handler的Mapping
*/

package controler

import (
	"common"
	"net/http"
)

type Handler struct {
	URL         string
	Permission  int
	HandlerFunc http.HandlerFunc
}

var (
	Handlers = []Handler{

		Handler{"/", common.Everyone, indexHandler},
		Handler{"/question/add", common.Authenticated, QuestionAddHandler},
		Handler{"/question/{id}", common.Everyone, QuestionViewHandler},
		Handler{"/questions", common.Everyone, QuestionSearchResultHandler},
		Handler{"/mypersonal", common.Authenticated, MyPersonalHandler},
		Handler{"/tags", common.Everyone, TagListHandler},
		Handler{"/tags/tagquestion", common.Everyone, TagQuestionHandler},
		Handler{"/member/{username}", common.Everyone, PersonalHandler},
		Handler{"/member/{username}/questions", common.Everyone, PersonalQuestionHandler},

		Handler{"/ajax/uploadimage", common.Authenticated, ImageUploadAjaxHandler},
		Handler{"/ajax/question", common.Everyone, QuestionAjaxHandler},
		Handler{"/ajax/answer", common.Everyone, AnswerAjaxHandler},
		Handler{"/ajax/tags", common.Everyone, TagAjaxHandler},
		Handler{"/ajax/user", common.Everyone, UserAjaxHandler},
		Handler{"/ajax/common", common.Everyone, CommonAjaxHandler},

		/*
			Handler{"/about", Everyone, staticHandler("about.html")},
			Handler{"/faq", Everyone, staticHandler("faq.html")},
			Handler{"/search", Everyone, searchHandler},
			Handler{"/admin", Administrator, adminHandler},
			Handler{"/admin/nodes", Administrator, adminListNodesHandler},
			Handler{"/admin/node/new", Administrator, adminNewNodeHandler},
			Handler{"/admin/site_categories", Administrator, adminListSiteCategoriesHandler},
			Handler{"/admin/site_category/new", Administrator, adminNewSiteCategoryHandler},
			Handler{"/admin/users", Administrator, adminListUsersHandler},
			Handler{"/admin/user/{userId}/activate", Administrator, adminActivateUserHandler},
			Handler{"/admin/article_categories", Administrator, adminListArticleCategoriesHandler},
			Handler{"/admin/article_category/new", Administrator, adminNewArticleCategoryHandler},
			Handler{"/admin/package_categories", Administrator, adminListPackageCategoriesHandler},
			Handler{"/admin/package_category/new", Administrator, adminNewPackageCategoryHandler},
			Handler{"/admin/package_category/{id}/edit", Administrator, adminEditPackageCategoryHandler},
			Handler{"/admin/link_exchanges", Administrator, adminListLinkExchangesHandler},
			Handler{"/admin/link_exchange/new", Administrator, adminNewLinkExchangeHandler},
			Handler{"/admin/link_exchange/{linkExchangeId}/edit", Administrator, adminEditLinkExchangeHandler},
			Handler{"/admin/link_exchange/{linkExchangeId}/delete", Administrator, adminDeleteLinkExchangeHandler},
			Handler{"/admin/ads", Administrator, adminListAdsHandler},
			Handler{"/admin/ad/new", Administrator, adminNewAdHandler},
			Handler{"/admin/ad/{id}/delete", Administrator, adminDeleteAdHandler},
			Handler{"/admin/ad/{id}/edit", Administrator, adminEditAdHandler},
		*/
		Handler{"/signup", common.Everyone, SignupHandler},
		Handler{"/signin", common.Everyone, SigninHandler},
		Handler{"/signout", common.Authenticated, SignoutHandler},
		Handler{"/profile", common.Authenticated, ProfileHandler},
		Handler{"/profile/photo", common.Authenticated, ChangePhotoHandler},
		Handler{"/profile/choose_default_photo", common.Authenticated, ChooseDefaultPhoto},
		/*			Handler{"/activate/{code}", Everyone, activateHandler},
					Handler{"/forgot_password", Everyone, forgotPasswordHandler},
					Handler{"/reset/{code}", Everyone, resetPasswordHandler},

					Handler{"/change_password", Authenticated, changePasswordHandler},

					Handler{"/profile/choose_default_avatar", Authenticated, chooseDefaultAvatar},

					Handler{"/nodes", Everyone, nodesHandler},
					Handler{"/go/{node}", Everyone, topicInNodeHandler},

					Handler{"/comment/{contentId}", Authenticated, commentHandler},
					Handler{"/comment/{commentId}/delete", Administrator, deleteCommentHandler},

					Handler{"/topics/latest", Everyone, latestTopicsHandler},
					Handler{"/topics/no_reply", Everyone, noReplyTopicsHandler},
					Handler{"/topic/new", Authenticated, newTopicHandler},
					Handler{"/new/{node}", Authenticated, newTopicHandler},
					Handler{"/t/{topicId}", Everyone, showTopicHandler},
					Handler{"/t/{topicId}/edit", Authenticated, editTopicHandler},
					Handler{"/t/{topicId}/delete", Administrator, deleteTopicHandler},

					Handler{"/member/{username}", Everyone, memberInfoHandler},
					Handler{"/member/{username}/topics", Everyone, memberTopicsHandler},
					Handler{"/member/{username}/replies", Everyone, memberRepliesHandler},
					Handler{"/follow/{username}", Authenticated, followHandler},
					Handler{"/unfollow/{username}", Authenticated, unfollowHandler},
					Handler{"/members", Everyone, membersHandler},
					Handler{"/members/all", Everyone, allMembersHandler},
					Handler{"/members/city/{cityName}", Everyone, membersInTheSameCityHandler},

					Handler{"/sites", Everyone, sitesHandler},
					Handler{"/site/new", Authenticated, newSiteHandler},
					Handler{"/site/{siteId}/edit", Authenticated, editSiteHandler},
					Handler{"/site/{siteId}/delete", Administrator, deleteSiteHandler},

					Handler{"/article/new", Authenticated, newArticleHandler},
					Handler{"/articles", Everyone, listArticlesHandler},
					Handler{"/a/{articleId}", Everyone, showArticleHandler},
					Handler{"/a/{articleId}/edit", Authenticated, editArticleHandler},

					Handler{"/packages", Everyone, packagesHandler},
					Handler{"/package/new", Authenticated, newPackageHandler},
					Handler{"/packages/{categoryId}", Everyone, listPackagesHandler},
					Handler{"/p/{packageId}", Everyone, showPackageHandler},
					Handler{"/p/{packageId}/edit", Authenticated, editPackageHandler},
					Handler{"/p/{packageId}/delete", Administrator, deletePackageHandler},

					Handler{"/download", Everyone, downloadHandler},
		*/
	}
)
