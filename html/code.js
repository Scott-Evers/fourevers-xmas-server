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

$(window).on("load", () => {
  loadData();

  $("#config").change(function () {
    selected_value = $("input[name='device']:checked").val();
    console.log("device radio button selected:", selected_value);
  });
});
