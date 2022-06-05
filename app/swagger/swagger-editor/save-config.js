function saveContent() {
    const content = window.localStorage.getItem('swagger-editor-content');
    console.log(content);
    let blob = new Blob([content], {type: "application/json;charset=utf-8"});

    // var data = new FormData();
    // data.append("upfile", new Blob(["CONTENT"], {type: "text/plain"}));
    // fetch("SERVER.SCRIPT", { method: "POST", body: blob });

    saveAs(blob, "swagger.json");
}

