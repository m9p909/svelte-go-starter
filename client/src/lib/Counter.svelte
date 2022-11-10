<script>
  import { serverURL } from "../env";

  let helloWorld = "";
  if (serverURL == undefined || serverURL == null) {
    console.log(import.meta.env);
    throw new Error("no base url found");
  }
  const sayHi = async () => {
    const output = await fetch(serverURL + "/api/helloWorld");
    helloWorld = await output.text();
  };
</script>

<button on:click={sayHi}>
  Message: {helloWorld}
</button>
