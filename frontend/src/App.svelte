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
            icon: '<svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_iconCarrier"><path d="M10.255 4.18806C9.84269 5.17755 8.68655 5.62456 7.71327 5.17535C6.10289 4.4321 4.4321 6.10289 5.17535 7.71327C5.62456 8.68655 5.17755 9.84269 4.18806 10.255C2.63693 10.9013 2.63693 13.0987 4.18806 13.745C5.17755 14.1573 5.62456 15.3135 5.17535 16.2867C4.4321 17.8971 6.10289 19.5679 7.71327 18.8246C8.68655 18.3754 9.84269 18.8224 10.255 19.8119C10.9013 21.3631 13.0987 21.3631 13.745 19.8119C14.1573 18.8224 15.3135 18.3754 16.2867 18.8246C17.8971 19.5679 19.5679 17.8971 18.8246 16.2867C18.3754 15.3135 18.8224 14.1573 19.8119 13.745C21.3631 13.0987 21.3631 10.9013 19.8119 10.255C18.8224 9.84269 18.3754 8.68655 18.8246 7.71327C19.5679 6.10289 17.8971 4.4321 16.2867 5.17535C15.3135 5.62456 14.1573 5.17755 13.745 4.18806C13.0987 2.63693 10.9013 2.63693 10.255 4.18806Z" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="paint-order: stroke; stroke: rgb(0, 0, 0);"></path><path d="M15 12C15 13.6569 13.6569 15 12 15C10.3431 15 9 13.6569 9 12C9 10.3431 10.3431 9 12 9C13.6569 9 15 10.3431 15 12Z" stroke-width="2" style="paint-order: stroke; stroke: rgb(0, 0, 0);"></path></g></svg>',
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
    let toolbar = vditor.vditor.options.toolbar
    for (let tool of toolbar) {
      tool.tipPosition = "s"
    }
    
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
