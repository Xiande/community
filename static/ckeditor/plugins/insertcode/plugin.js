CKEDITOR.plugins.add('insertcode', {
    requires: ['dialog'],
    init: function(a){
        var b = a.addCommand('insertcode', new CKEDITOR.dialogCommand('insertcode'));
        a.ui.addButton('insertcode', {
            label: "InsertCode",
            command: 'insertcode'
        });
        CKEDITOR.dialog.add('insertcode', this.path + 'dialogs/insertcode.js');
    }
});