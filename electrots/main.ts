// Declare vanilla js vars as type "any"
declare var require: any;
declare var __dirname: any;
declare var process: any;

const electron = require("electron");
const {app} = electron;
const {BrowserWindow} = electron;

// Keep a global ref of the window object to protect it from GC
let win;

function createWindow() {
  win = new BrowserWindow({width: 1080, height: 960});
  win.loadURL(`file://${__dirname}/index.html`);

  // Pop open devtools in case anything has gone awry on startup
  win.webContents.openDevTools();

  win.on("closed", () => {
    win = null;
  });
}

app.on("ready", createWindow);

app.on("window-all-closed", () => {
  if (process.platform !== "darwin") {
    app.quit();
  }
});

// Creates a new window when the icon is clicked on OSX dock
app.on("activate", () => {
  if (win === null) {
    createWindow();
  }
});
