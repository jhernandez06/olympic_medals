$(function () {
  get();
});

function get() {
  $("tbody").empty();
  var $tbody = $("#tbody");

  $.ajax({
    url: "/api/countries",
    type: "GET",
    crossDomain: true,
    dataType: "json",
    success: function (getApi) {
      $.each(getApi, function (idx, item) {
        var $tr = $("<tr></tr>");
        $tr.append(
          `<th scope="row"><div class="text-center"> ${item.Position} </div></th>`
        );
        $tr.append('<th scope="row"><div>' + item.Country + " </div></th>");
        $tr.append(
          '<th scope="row"><div class="text-center">' +
            item.GoldMedals +
            "</div></th>"
        );
        $tr.append(
          '<th scope="row"><div class="text-center"> ' +
            item.SilverMedals +
            " </div></th>"
        );
        $tr.append(
          '<th scope="row"><div class="text-center"> ' +
            item.BronzeMedals +
            " </div></th>"
        );
        $tr.append(
          '<th scope="row"><div class="text-center"> ' +
            item.AllMedals +
            " </div></th>"
        );

        $tbody.append($tr);
      });
    },
    error: function (request, msg, error) {
      alert("Error!");
    },
  });
}