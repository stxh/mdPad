<script>
  import { onMount } from "svelte"
  import Vditor from "vditor"
  import i18n from './assets/js/i18n'

  import { App } from "../bindings/mdPad"

  import * as wails from "@wailsio/runtime";
  
  let filename = "";
  let bChnaged = false;
  let aboutDialog;
  let vditor;

  function setWindowTitle() {
    let title = "mdPad"
    if (filename != "") {
      title += " - " + filename
    }

    if (bChnaged) {
      title += " *"
    }
    wails.Window.SetTitle(title)
    setSaveButtonStatus()
  }

  onMount(() => {
     // console.log("--> value=", value)
    vditor = new Vditor("vditor-container", {
        theme: "classic",
        cdn: "",
        toolbar: [
          {
            hotkey: '⌘n',
            name: 'new',
            tipPosition: 's',
            tip: $i18n.t('newFile'),
            className: 'right',
            icon: '<svg xmlns="http://www.w3.org/2000/svg" width="36" height="36" fill="currentColor" viewBox="2 2 20 20"><path fill-rule="evenodd" d="M9 2.221V7H4.221a2 2 0 0 1 .365-.5L8.5 2.586A2 2 0 0 1 9 2.22ZM11 2v5a2 2 0 0 1-2 2H4v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2h-7Z" clip-rule="evenodd"/></svg>',
            click () { clickNewFile(vditor) },
          },
          {
            hotkey: '⌘o',
            name: 'open',
            tipPosition: 's',
            tip: $i18n.t('openFile'),
            className: 'right',
            icon: '<svg aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="36" height="36" fill="currentColor" viewBox="2 2 20 20"><path fill-rule="evenodd" d="M4 4a2 2 0 0 0-2 2v12a2 2 0 0 0 .087.586l2.977-7.937A1 1 0 0 1 6 10h12V9a2 2 0 0 0-2-2h-4.532l-1.9-2.28A2 2 0 0 0 8.032 4H4Zm2.693 8H6.5l-3 8H18l3-8H6.693Z" clip-rule="evenodd"/></svg>',
            click () { clickOpenFile(vditor) },
          },
          {
            hotkey: '⌘s',
            name: 'save',
            tipPosition: 's',
            tip: $i18n.t('save'),
            className: 'right',
            icon: '<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" fill="currentColor" viewBox="2 2 20 20"><path fill-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V7.414A2 2 0 0 0 20.414 6L18 3.586A2 2 0 0 0 16.586 3H5Zm3 11a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v6H8v-6Zm1-7V5h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14 17h-4v-2h4v2Z" clip-rule="evenodd"/></svg>',
            click () { clickSaveFile(vditor) },
          },
          {
            hotkey: '⇧⌘s',
            name: 'save-as',
            tipPosition: 's',
            tip: $i18n.t('saveAs'),
            className: 'right',
            icon: '<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" fill="currentColor" viewBox="2 2 20 20"><path fill-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V7.414A2 2 0 0 0 20.414 6L18 3.586A2 2 0 0 0 16.586 3H5Zm10 11a3 3 0 1 1-6 0 3 3 0 0 1 6 0ZM8 7V5h8v2a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1Z" clip-rule="evenodd"/></svg>',
            click () { clickSaveAsFile(vditor) },
          },
          "|","undo","redo","|","emoji","headings","bold","italic","strike","link","|","list","ordered-list","check","outdent","indent",
          "|","quote","line","code","inline-code","insert-before","insert-after","|","table","edit-mode",
          "|",
          {
            hotkey: '⌘,',
            name: 'setup',
            tipPosition: 's',
            tip: $i18n.t('setup'),
            className: 'right',
            icon: '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M13.85 22.25h-3.7c-.74 0-1.36-.54-1.45-1.27l-.27-1.89c-.27-.14-.53-.29-.79-.46l-1.8.72c-.7.26-1.47-.03-1.81-.65L2.2 15.53c-.35-.66-.2-1.44.36-1.88l1.53-1.19c-.01-.15-.02-.3-.02-.46 0-.15.01-.31.02-.46l-1.52-1.19c-.59-.45-.74-1.26-.37-1.88l1.85-3.19c.34-.62 1.11-.9 1.79-.63l1.81.73c.26-.17.52-.32.78-.46l.27-1.91c.09-.7.71-1.25 1.44-1.25h3.7c.74 0 1.36.54 1.45 1.27l.27 1.89c.27.14.53.29.79.46l1.8-.72c.71-.26 1.48.03 1.82.65l1.84 3.18c.36.66.2 1.44-.36 1.88l-1.52 1.19c.01.15.02.3.02.46s-.01.31-.02.46l1.52 1.19c.56.45.72 1.23.37 1.86l-1.86 3.22c-.34.62-1.11.9-1.8.63l-1.8-.72c-.26.17-.52.32-.78.46l-.27 1.91c-.1.68-.72 1.22-1.46 1.22zm-3.23-2h2.76l.37-2.55.53-.22c.44-.18.88-.44 1.34-.78l.45-.34 2.38.96 1.38-2.4-2.03-1.58.07-.56c.03-.26.06-.51.06-.78s-.03-.53-.06-.78l-.07-.56 2.03-1.58-1.39-2.4-2.39.96-.45-.35c-.42-.32-.87-.58-1.33-.77l-.52-.22-.37-2.55h-2.76l-.37 2.55-.53.21c-.44.19-.88.44-1.34.79l-.45.33-2.38-.95-1.39 2.39 2.03 1.58-.07.56a7 7 0 0 0-.06.79c0 .26.02.53.06.78l.07.56-2.03 1.58 1.38 2.4 2.39-.96.45.35c.43.33.86.58 1.33.77l.53.22.38 2.55z"></path><circle cx="12" cy="12" r="3.5"></circle></svg>',
            click () {
              // let toolbar = vditor.vditor.options.toolbar
              // // disableToolbar(toolbar, ['save'])
              // console.log(toolbar)
              
              // const itemElement = toolbar['save'].children[0];
              // itemElement.classList.add("vditor-menu--disabled");
              
              // autolog.log(" file saved", "success", 2500);
            },
          },
          {
            hotkey: '⌘F1',
            name: 'about',
            tipPosition: 's',
            tip: $i18n.t('about'),
            className: 'right',
            icon: '<svg width="125px" height="125px" viewBox="0 0 24 24" fill="#000000" xmlns="http://www.w3.org/2000/svg" xmlns:bx="https://boxy-svg.com"><path d="M3 9.22843V14.7716C3 15.302 3.21071 15.8107 3.58579 16.1858L7.81421 20.4142C8.18929 20.7893 8.69799 21 9.22843 21H14.7716C15.302 21 15.8107 20.7893 16.1858 20.4142L20.4142 16.1858C20.7893 15.8107 21 15.302 21 14.7716V9.22843C21 8.69799 20.7893 8.18929 20.4142 7.81421L16.1858 3.58579C15.8107 3.21071 15.302 3 14.7716 3H9.22843C8.69799 3 8.18929 3.21071 7.81421 3.58579L3.58579 7.81421C3.21071 8.18929 3 8.69799 3 9.22843Z" stroke-linecap="round" stroke-linejoin="round" style="fill-opacity: 0; paint-order: fill; stroke-width: 2px; stroke: rgb(0, 0, 0);"></path><path d="M12 8V13" stroke="#323232" stroke-width="2" stroke-linecap="round"></path><path d="M12 16V15.9888" stroke="#323232" stroke-width="2" stroke-linecap="round"></path></svg>',
            click () {showAboutDialog()},
          },
      ],
      toolbarConfig: {pin: true, hide: true},
      input:((_) => {
        // console.log(value)
        bChnaged = true
        setWindowTitle()
      }),
      after:(() => {
        wails.Window.SetMinSize(1024, 178)
        // console.log(vditor)
        initMd(vditor)
      }),
    })

    // console.log(vditor)
    
  })
  
  function initMd(vditor) {
    // 修复toolbar的tip弹出方向
    const toolButtons = document.querySelectorAll('div.vditor-toolbar button');

    toolButtons.forEach(button => {
      button.classList.remove("vditor-tooltipped__ne")
      button.classList.remove("vditor-tooltipped__nw")
      button.classList.remove("vditor-tooltipped__n")
      button.classList.add("vditor-tooltipped__s")
    })
    
    App.GetMdContent().then((result) => {
      // console.log("result ->", result)
      filename = result.filename;
      setWindowTitle()
      vditor.setValue(result.value)
      vditor.updateToolbarConfig({hide: false})
    });
  }

  function clickNewFile(vditor) {
    App.GoNewFile() .then(() => {
      filename = "";
      setWindowTitle()
      vditor.setValue("")
    })
  }

  function clickOpenFile(vditor) {
      App.GoOpenFile().then((result) => {
        //  console.log("result ->", result.filename, result.value)
        filename = result.filename;
        bChnaged = false;
        setWindowTitle()
        vditor.setValue(result.value, true)
      }).catch((error) => {
        console.error("Read File", error); // 处理失败的情况
      });
  }

  function clickSaveFile(vditor) {
    App.GoSaveFile(vditor.getValue()).then((result) => {
      bChnaged = false
      setWindowTitle()
      autolog.log($i18n.t("saveToast") + filename + " " + result, "success")
    })
  }

  function clickSaveAsFile(vditor) {
    App.GoSaveAsFile(vditor.getValue()).then((result) => {
      // console.log("result", result, "file", file) 
      filename = result.filename
      bChnaged = false
      setWindowTitle()
      autolog.log($i18n.t("saveToast") + filename + " " + result.value, "success")
    })
  }

  function setSaveButtonStatus() {
    let saveButton = document.querySelector('[data-type="save"]');
    if (bChnaged && filename)
      saveButton?.classList.remove("vditor-menu--disabled")
    else
      saveButton?.classList.add("vditor-menu--disabled")
  }

  // About Dialog Functionality
  function showAboutDialog() {
    aboutDialog.classList.add("active");
  }

  function hideAboutDialog() {
    aboutDialog.classList.remove("active");
  }
</script>


<main>
  <div id="contains">
    <div id="vditor-container"></div>

    <!-- About Dialog -->
    <div id="about-dialog" class="dialog-overlay" bind:this={aboutDialog} >
      <div class="dialog-content">
          <h2 data-i18n="aboutTitle">About</h2>
          <p data-i18n="aboutContent">
              mdPad is a desktop Markdown application developed
              using the wails3 framework. It combines the powerful
              backend capabilities of go with modern web frontend svelte and js vditor technologies to provide a smooth and
              efficient Markdown editing experience.
          </p>
          <button id="close-about" data-i18n="closeButton" on:click={hideAboutDialog}>
              Close
          </button>
      </div>
    </div>
  
  </div>
</main>
