<script>
  import { onMount } from "svelte"
  import Vditor from "vditor"

  import * as app from "../wailsjs/go/main/App";
  import { WindowSetTitle } from "../wailsjs/runtime/runtime.js";
  import * as runtime from "../wailsjs/runtime/runtime.js";

  import { SvelteToast, toast } from '@zerodevx/svelte-toast';

  function showToastMessage(message) {
      toast.push(message);
  }

  let filename = "";

  function setWindowTitle() {
    let title = "mdPad"
    if (filename != "") {
      title += " - " + filename
    }
    // console.log("--> setWindowTitle title=", title)
    WindowSetTitle(title)
  }

  onMount(() => {
     // console.log("--> value=", value)
    const vditor = new Vditor("vditor-container", {
        theme: "classic",
        toolbar: [
          {
            hotkey: '^n',
            name: 'new',
            tipPosition: 's',
            tip: '新建',
            className: 'right',
            icon: '<svg xmlns="http://www.w3.org/2000/svg" width="36" height="36" fill="currentColor" viewBox="2 2 20 20"><path fill-rule="evenodd" d="M9 2.221V7H4.221a2 2 0 0 1 .365-.5L8.5 2.586A2 2 0 0 1 9 2.22ZM11 2v5a2 2 0 0 1-2 2H4v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2h-7Z" clip-rule="evenodd"/></svg>',
            click () { clickNewFile(vditor) },
          },
          {
            hotkey: '^o',
            name: 'open',
            tipPosition: 's',
            tip: '打开文件',
            className: 'right',
            icon: '<svg aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="36" height="36" fill="currentColor" viewBox="2 2 20 20"><path fill-rule="evenodd" d="M4 4a2 2 0 0 0-2 2v12a2 2 0 0 0 .087.586l2.977-7.937A1 1 0 0 1 6 10h12V9a2 2 0 0 0-2-2h-4.532l-1.9-2.28A2 2 0 0 0 8.032 4H4Zm2.693 8H6.5l-3 8H18l3-8H6.693Z" clip-rule="evenodd"/></svg>',
            click () { clickOpenFile(vditor) },
          },
          {
            hotkey: '^S',
            name: 'save',
            tipPosition: 's',
            tip: '保存',
            className: 'right',
            icon: '<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" fill="currentColor" viewBox="2 2 20 20"><path fill-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V7.414A2 2 0 0 0 20.414 6L18 3.586A2 2 0 0 0 16.586 3H5Zm3 11a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v6H8v-6Zm1-7V5h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14 17h-4v-2h4v2Z" clip-rule="evenodd"/></svg>',
            click () { clickSaveFile(vditor) },
          },
          {
            hotkey: '^⇧s',
            name: 'save as',
            tipPosition: 's',
            tip: '另存为...',
            className: 'right',
            icon: '<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" fill="currentColor" viewBox="2 2 20 20"><path fill-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V7.414A2 2 0 0 0 20.414 6L18 3.586A2 2 0 0 0 16.586 3H5Zm10 11a3 3 0 1 1-6 0 3 3 0 0 1 6 0ZM8 7V5h8v2a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1Z" clip-rule="evenodd"/></svg>',
            click () { clickSaveAsFile(vditor) },
          },
          "|","undo","redo","|","emoji","headings","bold","italic","strike","link","|","list","ordered-list","check","outdent","indent",
          "|","quote","line","code","inline-code","insert-before","insert-after","|","table","content-theme",
      ],
      toolbarConfig: {pin: true},
      after:(() => {initMd(vditor)}),
    })

  })

  function initMd(vditor) {
    app.GetMdContent().then((result) => {
      // console.log("result ->", result)
      filename = result.filename;
      setWindowTitle()
      vditor.setValue(result.value)
    });
  }

  function clickNewFile(vditor) {
    app.GoNewFile() .then(() => {
      filename = "";
      setWindowTitle()
      vditor.setValue("")
    })
  }

  function clickOpenFile(vditor) {
    app.GoOpenFile().then((result) => {
      //  console.log("result ->", result.filename, result.value)
      filename = result.filename;
      setWindowTitle()
      vditor.setValue(result.value, true)
    })
  }

  function clickSaveFile(vditor) {
    app.GoSaveFile(vditor.getValue()).then((result) => {
      showToastMessage("Save to file: "+ filename + " " + result)
    })
  }

  function clickSaveAsFile(vditor) {
    app.GoSaveAsFile(vditor.getValue()).then((result, file) => {
      // console.log("result", result, "file", file) 
      filename = file;
      showToastMessage("Save to file: "+ filename + " " + result)
    })
  }

</script>


<main>
  <div id="vditor-container"></div>
</main>

<SvelteToast />
