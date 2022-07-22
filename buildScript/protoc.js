const os = require("os");
const child = require("child_process");
const resolve = require('path').resolve;

const configure = require("../buildConfig/protoc.json");



configure.basePath = resolve(configure.basePath);
configure.outputRootPath = resolve(configure.outputRootPath);



const commandTemplateText = `protoc -I=${configure.basePath} --$OUTPUT_LANG$_out=$OUTPUT_DIR$ $SOURCE_FILES$`;



function Target(lang,output) {
    this.lang = lang;
    this.output = output;
}


const readAllFiles = (base) => {
    let sh = "";
    let shName = "bash"; 
    if("win32" == os.platform()) {
        sh = `Get-ChildItem -Path ${base} -Include *.proto -Depth 0 -Recurse | Select-Object  -Unique  | ForEach-Object {$_.FullName}`;
        shName = "powershell.exe";
    }
    let out = child.execSync(sh,{"shell":shName});
    return out.toString().split("\n").filter(value => value != "");
};

const makeProtoCommand = (target,srcFiles) => {
    return commandTemplateText.replace("$OUTPUT_LANG$",target.lang)
    .replace("$OUTPUT_DIR$",resolve(configure.outputRootPath,target.output))
    .replace("$SOURCE_FILES$",srcFiles);
}

const runProtocProcess = (target,fileName) => {
    let command = makeProtoCommand(target,fileName);
    console.log(command);
    try {
        child.execSync(command);
    }catch(e) {
        if(e.stdout.length != 0)console.log(e.stdout.toString());
        if(e.stderr.length != 0)console.error("err : " + e.stderr.toString());
    }
};

const trans = (target) => {
    let files = readAllFiles(configure.basePath);
    let t = new Target(target.lang,target.output);

    files.forEach(value => runProtocProcess(t,value));

};


const main = () => {
    configure.target.forEach(trans);
};


main();