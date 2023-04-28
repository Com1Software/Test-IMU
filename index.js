import { wasmBrowserInstantiate } from "/demo-util/instantiateWasm.js";

const go = new Go(); // Defined in wasm_exec.js. Don't forget to add this in your index.html.

const runWasmAdd = async () => {

  const importObject = go.importObject;

  const wasmModule = await wasmBrowserInstantiate("./main.wasm", importObject);

  go.run(wasmModule.instance);

  const addResult = wasmModule.instance.exports.add(24, 24);

  document.body.textContent = `Hello World! addResult: ${addResult}`;
};
runWasmAdd();
