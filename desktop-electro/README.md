Here I am combining:
* [Electron](http://electron.atom.io/)
* [Ionic 2 Framework](http://ionic.io/)

And writing all code in [TypeScript](http://www.typescriptlang.org/).

The idea is to see how feasible it is to write an isomorphic, strongly typed desktop app using a pretty web framework and being able to drop the same front-end into Ionic's flavor of [cordova](https://cordova.apache.org/) and make a mobile app out of it for iOS/Android.

Disclaimer: this might not work anywhere but my laptop.

Instructions:
```sh
# You will need the Electron prebuilt binary, and gulp. If needed:
npm i -g electron-prebuilt
# in desktop-electro/electro:
npm install
# then:
gulp build
# then, in desktop-electro:
electron .
```
