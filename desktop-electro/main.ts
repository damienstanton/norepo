// Delcare the vanilla js vars as "any" type
declare var require: any;
declare var __dirname: any;
declare var process: any;

const electron = require("electron");
const {app} = electron;
const {BrowserWindow} = electron;

let win;

function createWindow() {
  win = new BrowserWindow({width: 800, height: 600});
  win.loadURL(`file://${__dirname}/electro/www/index.html`);

  // DEBUG open dev tools in case there is a problem
  // win.webContents.openDevTools();

  win.on("closed", () => {
    win = null;
  });
}

app.on("ready", createWindow);

// The following two functions ensure good behavior on OSX
app.on("window-all-closed", () => {
  if (process.platform !== "darwin") {
    app.quit();
  }
});

app.on("activate", () => {
  if (win === null) {
    createWindow();
  }
});
