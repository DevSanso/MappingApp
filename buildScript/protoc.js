const os = require("os");
const child = require("child_process");
const resolve = require('path').resolve;
const parse = require('path').parse;
const readdirSync = require('fs').readdirSync;
const configure = require("../buildConfig/protoc.json");



configure.basePath = resolve(configure.basePath);
configure.outputRootPath = resolve(configure.outputRootPath);



//const commandTemplateText = `protoc -I=${configure.basePath} --$OUTPUT_LANG$_out=$OUTPUT_DIR$ $SOURCE_FILES$`;
const commandTemplateText = `protoc -I=$BASE_DIR$ --$OUTPUT_LANG$_out=$OUTPUT_DIR$ $SOURCE_FILES$`;


function Target(lang,output) {
    this.lang = lang;
    this.output = output;
}
const readEntryFilePath = (base,entryRequest,entryResponse) => {
    let sh = "";
    let shName = "bash"; 
    if("win32" == os.platform()) {
        sh = `Get-ChildItem -Path ${base} -Include ${entryRequest},${entryResponse} -Depth 0 -Recurse | Select-Object  -Unique  | ForEach-Object {$_.FullName}`;
        shName = "powershell.exe";
    }
    let out = child.execSync(sh,{"shell":shName});
    return out.toString().split("\n").filter(value => value != "");
};

const makeSrcFileStr = (srcFiles) => {
    return srcFiles.reduce((pre,value)=> pre + `${value} `,"");
};

const makeProtoCommand = (target,base,srcFiles) => {
    return commandTemplateText.replace("$OUTPUT_LANG$",target.lang)
    .replace("$OUTPUT_DIR$",resolve(configure.outputRootPath,target.output))
    .replace("$BASE_DIR$",base)
    .replace("$SOURCE_FILES$",makeSrcFileStr(srcFiles));
};

const runProtocProcess = (target,base,files) => {
    let command = makeProtoCommand(target,base,files);
    console.log(command);
    try {
        child.execSync(command);
    }catch(e) {
        if(e.stdout.length != 0)console.log(e.stdout.toString());
        if(e.stderr.length != 0)console.error("err : " + e.stderr.toString());
    }
};

const getEntryFileParentDir = (path) => {
    const p = parse(path);
    return p.dir;
};

const readFilesFromPath = (path) => {
    const list = readdirSync(path, { withFileTypes: true});
    const dirents = list
    .filter((dirent) => dirent.isFile())
    .filter(dirent => parse(dirent.name).ext == ".proto");

    return dirents.map(value => value.name);
};

const trans = (target) => {
    let entryFile = readEntryFilePath(configure.basePath,
        configure.entrypoint.request,
        configure.entrypoint.reeponse);
    
    const srcs = entryFile.map(value => {
        const dir = getEntryFileParentDir(value);
        const files = readFilesFromPath(dir);
        return {
            dir : dir,
            files : files
        }
    });


    let t = new Target(target.lang,target.output);
    
    srcs.forEach(value => runProtocProcess(t,value.dir,value.files));

};


const main = () => {
    configure.target.forEach(trans);
};


main();