//Search
$("#search").submit(function (e) {
    e.preventDefault();
    var form = $(this).serialize();
    var req = "/search/?" + form;
    $.ajax({
        type: "GET",
        url: req,
        data: form,
        success: function (data) {
            $("#resrow").load(req + " #resrow");
        }
    });
});
//ekleModal
function ekle(e) {

    var img = $(e).data("img");
    var price = $(e).data("price");
    var title = $(e).data("title");
    var publisher = $(e).data("publisher");
    var author = $(e).data("author");
    var modal = $("#ekleModal");

    modal.modal();
    modal.find("img").attr("src", img);
    modal.find(".price").val(price);
    modal.find(".title").val(title);
    modal.find(".publisher").val(publisher);
    modal.find(".author").val(author);
}