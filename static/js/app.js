var getCode = function () {
  return $("#code").val();
};

$("#run").on("click", function (event) {
  $.ajax({
    url: "/code",
    data: getCode(),
    success: function (data, textStatus, jqXHR) {
      console.log(data);
    }
  })
});

