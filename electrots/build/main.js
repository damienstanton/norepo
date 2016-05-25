var electron = require("electron");
var app = electron.app;
var BrowserWindow = electron.BrowserWindow;
var win;
function createWindow() {
    win = new BrowserWindow({ width: 1080, height: 960 });
    win.loadURL("file://" + __dirname + "/index.html");
    win.webContents.openDevTools();
    win.on("closed", function () {
        win = null;
    });
}
app.on("ready", createWindow);
app.on("window-all-closed", function () {
    if (process.platform !== "darwin") {
        app.quit();
    }
});
app.on("activate", function () {
    if (win === null) {
        createWindow();
    }
});
//# sourceMappingURL=main.js.map