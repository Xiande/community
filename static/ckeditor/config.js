/**
 * @license Copyright (c) 2003-2015, CKSource - Frederico Knabben. All rights reserved.
 * For licensing, see LICENSE.md or http://ckeditor.com/license
 */

CKEDITOR.editorConfig = function( config ) {
	// Define changes to default configuration here. For example:
	 config.language = CCConfig.SiteLanguage;
    // config.uiColor = '#AADC6E';
    config.image_previewText = ' ';
    config.extraPlugins = "dragdrop,imagepaste,selfpreview";
    config.toolbar = 'Full';
    //config.allowedContent = true;
    
	config.toolbar_Full = [
        ['Bold', 'Italic', 'Underline', 'Strike', '-', 'Subscript', 'Superscript'],
        ['NumberedList', 'BulletedList', '-', 'Outdent', 'Indent', 'Blockquote'],
        ['JustifyLeft', 'JustifyCenter', 'JustifyRight', 'JustifyBlock'],
        ['Link', 'Unlink', 'Anchor'],
        ['Image', 'Table', 'HorizontalRule', 'Smiley', 'SpecialChar', 'PageBreak'],
        ['TextColor', 'BGColor'], [ 'Maximize','selfpreview']
    ];    
	//'Preview',
	/*
    config.toolbarGroups = [
		{ name: 'clipboard',   groups: [ 'clipboard', 'undo' ] },
		{ name: 'editing',     groups: [ 'find', 'selection', 'spellchecker' ] },
		{ name: 'links' },
		{ name: 'insert' },
		{ name: 'forms' },
		{ name: 'tools' },
		{ name: 'document',	   groups: [ 'mode', 'document', 'doctools' ] },
		{ name: 'others' },
		'/',
		{ name: 'basicstyles', groups: [ 'basicstyles', 'cleanup' ] },
		{ name: 'paragraph',   groups: ['list', 'indent', 'blocks', 'align', 'bidi' ] },
		{ name: 'styles' },
		{ name: 'colors' },
		{ name: 'about' }
	];
	*/
    config.filebrowserImageUploadUrl = CCConfig.L_Menu_BaseUrl + "/ajax/uploadimage";
    // Remove some buttons, provided by the standard plugins, which we don't
	// need to have in the Standard(s) toolbar.
	config.removeButtons = 'Underline,Subscript,Superscript';

	// Se the most common block elements.
	config.format_tags = 'p;h1;h2;h3;pre';

	// Make dialogs simpler.
	config.removeDialogTabs = 'image:advanced;link:advanced';
	
	// Custom Skin
	config.skin='themepixels';
	config.browserContextMenuOnCtrl = true;
	config.dialog_backgroundCoverColor = '#000' ;
	config.dialog_backgroundCoverOpacity = 0.65;
	config.dragdropConfig = {
	      backend: 'imgur',
	      settings: {
	          clientId: 'YourImgurClientID'
	      }
  	}
};
