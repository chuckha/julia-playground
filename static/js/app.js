var getCode = function () {
  return $("#code").val();
};

$("#run").on("click", function (event) {
  $.ajax({
    type: "POST",
    url: "/code",
    data: {
      "code": getCode()
    },
    success: function (data, textStatus, jqXHR) {
      console.log(data);
    }
  })
});

