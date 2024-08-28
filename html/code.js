const loadDevices = () => {
  $.get("/devices", (data) => {
    console.log(data);
    const devices = JSON.parse(data);
    devices.forEach((dev) => {
      let devcontrols = $("#devices")
        .first()
        .children()
        .append(
          `<div><input type="radio" name="device" id="${dev.ID}" value="${dev.ID}"><label for="${dev.ID}">${dev.Name}</label></div>`
        );
    });
  });
};

const loadData = () => {
  loadDevices();
};

const updateData = () => {
  const config = {
    DeviceID: $("input[name='device']:checked").val(),
  };
  console.log("updateData", config);
  $.ajax({
    url: "/devices",
    type: "PUT",
    data: JSON.stringify(config),
    contentType: "application/json",
    success: function (response) {
      console.info("Config updated");
    },
    error: function (error) {
      console.error("Error updating config", error);
    },
  });
};

$(window).on("load", () => {
  loadData();

  let selected_value = "";
  $("#config").change(function () {
    selected_value = $("input[name='device']:checked").val();
    console.log("device radio button selected:", selected_value);
  });

  $("#submit_button").click(updateData);
});
