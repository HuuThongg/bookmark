// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
  namespace App {
    // interface Error {}
    // interface Locals {}
    // interface PageData {}
    // interface PageState {}
    // interface Platform {}
  }
}

declare namespace globalThis {
  var handleCredentialResponse: (response: Object) => void;
}
export { };

// declare namespace App { }
//
// declare namespace globalThis {
//   var handleToken: (response: Object) => void;
// }
//
// declare namespace globalThis {
//   var handleCredentialResponse: (response: Object) => void;
// }
