/**
 * Created by pcrbrandao on 30/07/17.
 */
$(function() {
    $("#loginForm").submit(validateLogin)
});

function validateLogin(event) {
    event.preventDefault();

    var $form = $( this ),
        url = $form.attr("action");

    var posting = $.post( url, $form.serialize());
    posting.done(function(data){
        console.log(data);
        document.write(data);
    });
    posting.fail(function(data) {
        console.log(data)
       $("#mensagem_erro").text(data.responseText);
    });
}