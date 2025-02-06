import{bL as T,I as v,w as F,bM as Y,bN as l,f as g,bt as w,bP as P,b0 as _,bO as N,bS as k,bT as V,J as H,U as W,bh as D,bR as O,bV as A,bu as j}from"./vendor-CdkM_rDc.js";import{_ as $,A as G}from"./image-resize-BDYvEVDc.js";const J={class:"tool-settings"},K={class:"dimension-inputs"},Q={class:"dimensions-row"},Z={class:"dimensions-inputs"},ee={class:"input-fields"},te={class:"aspect-ratio-select"},ie={class:"crop-position"},oe={class:"position-inputs"},le={class:"input-group"},ne={class:"input-group"},ae={class:"action-buttons"},se=["disabled"],de=["disabled"],re=["disabled"],ue=T({__name:"CropToolsPanel",props:{width:{},height:{},positionX:{},positionY:{},disabled:{type:Boolean}},emits:["update:width","update:height","update:positionX","update:positionY","resize-image","crop","reset"],setup(B,{emit:E}){const c=B,r=E,C=[{title:"FreeForm",value:"freeform"},{title:"Original",value:"original"},{title:"Custom",value:"custom"},{title:"1:1 (Square)",value:"1:1"},{title:"4:3 (Monitor)",value:"4:3"},{title:"14:9",value:"14:9"},{title:"16:9 (Widescreen)",value:"16:9"},{title:"16:10",value:"16:10"},{title:"2:1",value:"2:1"},{title:"3:1 (Panaroma)",value:"3:1"},{title:"4:1",value:"4:1"},{title:"3:2 (35mm film)",value:"3:2"},{title:"5:4",value:"5:4"},{title:"7:5",value:"7:5"},{title:"19:10",value:"19:10"},{title:"21:9 (Cinemascope)",value:"21:9"},{title:"32:9 (Super Ultra Wide)",value:"32:9"},{title:"Facebook profile 170x170",value:"170:170"},{title:"Facebook cover 820x312",value:"820:312"},{title:"Facebook post 1200x900",value:"1200:900"},{title:"Facebook ad 1280x720",value:"1280:720"},{title:"Instagram profile 110x110",value:"110:110"},{title:"Instagram post 320x320",value:"320:320"},{title:"Instagram story 1080x1920",value:"1080:1920"},{title:"Twitter profile 400x400",value:"400:400"},{title:"Twitter header 1500x1500",value:"1500:1500"},{title:"Twitter image 1024x512",value:"1024:512"},{title:"Twitter card 1200x628",value:"1200:628"},{title:"Twitter ad 1200x675",value:"1200:675"},{title:"Youtube profile 800x800",value:"800:800"},{title:"Youtube channel art 2560x1440",value:"2560:1440"},{title:"Youtube thumb 1280x720",value:"1280:720"},{title:"Web mini 1024x768",value:"1024:768"},{title:"Web small 1280x800",value:"1280:800"},{title:"Web common 1366x768",value:"1366:768"},{title:"Web medium 1440x900",value:"1440:900"},{title:"Fll HD 1920x1080",value:"1920:1080"},{title:"Ultra HD 4x 3840x2160",value:"3840:2160"},{title:"Paper A4 2480x3508",value:"2480:3508"},{title:"Paper A5 1748x2480",value:"1748:2480"},{title:"Paper A6 1280x1748",value:"1280:1748"},{title:"Paper letter 2400x3300",value:"2400:3300"}],h=v("freeform"),m=F(()=>c.width),y=F(()=>c.height),b=F(()=>c.positionX),f=F(()=>c.positionY),S=F(()=>{const i=parseInt(c.width),e=parseInt(c.height);return i>0&&e>0}),z=i=>{if(i==="freeform"||i==="original")return;const[e,n]=i.split(":").map(Number);e&&n&&(r("update:width",e.toString()),r("update:height",n.toString()))},u=(i,e)=>{const n=parseFloat(i)||0;switch(e){case"width":r("update:width",n.toString());break;case"height":r("update:height",n.toString());break;case"positionX":r("update:positionX",n.toString());break;case"positionY":r("update:positionY",n.toString());break}},x=i=>{const e=i.target;e.value=e.value.replace(/-/g,"")},U=()=>{r("resize-image",{width:parseInt(c.width),height:parseInt(c.height)})},o=()=>{r("crop")},t=()=>{r("reset")};return(i,e)=>{const n=k("v-text-field"),s=k("v-select"),p=k("v-icon");return V(),Y("div",{class:N(["tools-panel",{"tools-panel--disabled":i.disabled}])},[l("div",J,[e[16]||(e[16]=l("h2",{class:"settings-title"},"Crop Rectangle",-1)),l("div",K,[l("div",Q,[e[9]||(e[9]=l("div",{class:"dimensions-header"},[l("div",{class:"dimensions-labels"},[l("label",null,"Width"),l("label",null,"Height")])],-1)),l("div",Z,[l("div",ee,[g(n,{"model-value":m.value,disabled:i.disabled,"onUpdate:modelValue":e[0]||(e[0]=d=>u(d,"width")),onInput:x,type:"number",min:"0",density:"compact",variant:"outlined","hide-details":"",class:"size-input",onWheel:e[1]||(e[1]=w(()=>{},["prevent"]))},null,8,["model-value","disabled"]),g(n,{"model-value":y.value,disabled:i.disabled,"onUpdate:modelValue":e[2]||(e[2]=d=>u(d,"height")),onInput:x,type:"number",min:"0",density:"compact",variant:"outlined","hide-details":"",class:"size-input",onWheel:e[3]||(e[3]=w(()=>{},["prevent"]))},null,8,["model-value","disabled"])])])])]),l("div",te,[e[10]||(e[10]=l("label",null,"Aspect Ratio",-1)),g(s,{modelValue:h.value,"onUpdate:modelValue":[e[4]||(e[4]=d=>h.value=d),z],items:C,"item-title":"title","item-value":"value",disabled:i.disabled,density:"compact",variant:"outlined","hide-details":"",class:"ratio-select"},null,8,["modelValue","disabled"])]),l("div",ie,[e[13]||(e[13]=l("h3",{class:"section-title"},"Crop Position",-1)),l("div",oe,[l("div",le,[e[11]||(e[11]=l("label",null,"Position (X)",-1)),g(n,{"model-value":b.value,disabled:i.disabled,"onUpdate:modelValue":e[5]||(e[5]=d=>u(d,"positionX")),onInput:x,type:"number",density:"compact",variant:"outlined","hide-details":"",class:"position-input",onWheel:e[6]||(e[6]=w(()=>{},["prevent"]))},null,8,["model-value","disabled"])]),l("div",ne,[e[12]||(e[12]=l("label",null,"Position (Y)",-1)),g(n,{"model-value":f.value,disabled:i.disabled,"onUpdate:modelValue":e[7]||(e[7]=d=>u(d,"positionY")),onInput:x,type:"number",density:"compact",variant:"outlined","hide-details":"",class:"position-input",onWheel:e[8]||(e[8]=w(()=>{},["prevent"]))},null,8,["model-value","disabled"])])])]),l("div",ae,[l("button",{class:"crop-btn",onClick:o,disabled:i.disabled},[g(p,{size:"small",class:"me-1"},{default:P(()=>e[14]||(e[14]=[_("mdi-crop")])),_:1}),e[15]||(e[15]=_(" Crop "))],8,se),l("button",{class:"reset-btn",onClick:t,disabled:i.disabled}," Reset ",8,de)]),l("button",{class:"resize-btn",onClick:U,disabled:!S.value||i.disabled}," Resize Image → ",8,re)])],2)}}}),ce=$(ue,[["__scopeId","data-v-ecc51f3c"]]),pe={class:"canvas-container"},ve={key:0,class:"image-container"},me=["src"],ge=T({__name:"CropPreviewPanel",props:{imageElement:{}},emits:["fileUpload","imageClick"],setup(B,{emit:E}){const c=B,r=v(null),C=v(null),h=v({x:0,y:0,width:0,height:0}),m=v({isDragging:!1,startX:0,startY:0,handle:""}),y=o=>{var L;if(!c.imageElement)return o;const t=(L=r.value)==null?void 0:L.querySelector(".image-container");if(!t)return o;const i=t.querySelector("img");if(!i)return o;const e=i.getBoundingClientRect(),n=t.getBoundingClientRect(),s=e.left-n.left,p=e.top-n.top,d=s+e.width,a=p+e.height,R=20,I=Math.max(o.width,R),M=Math.max(o.height,R),X=Math.min(Math.max(o.x,s),d-I),q=Math.min(Math.max(o.y,p),a-M);return{x:X,y:q,width:I,height:M}},b=o=>{if(!m.value.isDragging)return;const t=o.clientX-m.value.startX,i=o.clientY-m.value.startY;let e={...h.value};switch(m.value.handle){case"move":e={...e,x:e.x+t,y:e.y+i};break;case"nw":e={...e,x:e.x+t,y:e.y+i,width:e.width-t,height:e.height-i};break;case"n":e={...e,y:e.y+i,height:e.height-i};break;case"ne":e={...e,y:e.y+i,width:e.width+t,height:e.height-i};break;case"w":e={...e,x:e.x+t,width:e.width-t};break;case"e":e={...e,width:e.width+t};break;case"sw":e={...e,x:e.x+t,width:e.width-t,height:e.height+i};break;case"s":e={...e,height:e.height+i};break;case"se":e={...e,width:e.width+t,height:e.height+i};break}h.value=y(e),m.value.startX=o.clientX,m.value.startY=o.clientY},f=o=>{m.value.isDragging&&requestAnimationFrame(()=>{b(o)})},S=()=>{var R;if(!c.imageElement)return;const o=(R=r.value)==null?void 0:R.querySelector(".image-container");if(!o)return;const t=o.querySelector("img");if(!t)return;const i=t.getBoundingClientRect(),e=o.getBoundingClientRect(),n=.8,s=i.width*n,p=i.height*n,d=i.left-e.left+(i.width-s)/2,a=i.top-e.top+(i.height-p)/2;h.value={x:d,y:a,width:s,height:p}},z=()=>{requestAnimationFrame(()=>{S()})},u=(o,t)=>{m.value={isDragging:!0,startX:o.clientX,startY:o.clientY,handle:t}},x=()=>{m.value.isDragging=!1},U=()=>{var o;(o=C.value)==null||o.click()};return H(()=>c.imageElement,o=>{o&&S()}),W(()=>{window.addEventListener("mousemove",f),window.addEventListener("mouseup",x)}),D(()=>{window.removeEventListener("mousemove",f),window.removeEventListener("mouseup",x)}),(o,t)=>(V(),Y("div",{class:"preview-panel",ref_key:"previewPanelRef",ref:r},[l("div",pe,[o.imageElement?(V(),Y("div",ve,[l("img",{src:o.imageElement.src,style:{maxWidth:"100%",maxHeight:"100%",objectFit:"contain"},draggable:"false",onLoad:z},null,40,me),l("div",{class:"crop-box",style:O({left:`${h.value.x}px`,top:`${h.value.y}px`,width:`${h.value.width}px`,height:`${h.value.height}px`}),onMousedown:t[8]||(t[8]=i=>u(i,"move"))},[t[10]||(t[10]=A('<div class="grid-lines" data-v-36feed0c><div class="grid-line vertical" style="left:33.33%;" data-v-36feed0c></div><div class="grid-line vertical" style="left:66.66%;" data-v-36feed0c></div><div class="grid-line horizontal" style="top:33.33%;" data-v-36feed0c></div><div class="grid-line horizontal" style="top:66.66%;" data-v-36feed0c></div></div>',1)),l("div",{class:"handle nw",onMousedown:t[0]||(t[0]=w(i=>u(i,"nw"),["stop"]))},null,32),l("div",{class:"handle n",onMousedown:t[1]||(t[1]=w(i=>u(i,"n"),["stop"]))},null,32),l("div",{class:"handle ne",onMousedown:t[2]||(t[2]=w(i=>u(i,"ne"),["stop"]))},null,32),l("div",{class:"handle w",onMousedown:t[3]||(t[3]=w(i=>u(i,"w"),["stop"]))},null,32),l("div",{class:"handle e",onMousedown:t[4]||(t[4]=w(i=>u(i,"e"),["stop"]))},null,32),l("div",{class:"handle sw",onMousedown:t[5]||(t[5]=w(i=>u(i,"sw"),["stop"]))},null,32),l("div",{class:"handle s",onMousedown:t[6]||(t[6]=w(i=>u(i,"s"),["stop"]))},null,32),l("div",{class:"handle se",onMousedown:t[7]||(t[7]=w(i=>u(i,"se"),["stop"]))},null,32)],36)])):(V(),Y("div",{key:1,class:"upload-placeholder",onClick:U},[l("input",{type:"file",ref_key:"fileInput",ref:C,onChange:t[9]||(t[9]=i=>o.$emit("fileUpload",i)),accept:"image/*",style:{display:"none"}},null,544),t[11]||(t[11]=A('<div class="upload-content" data-v-36feed0c><svg class="upload-icon" viewBox="0 0 1024 1024" fill="currentColor" data-v-36feed0c><path d="M200.32 318.506667a324.906667 324.906667 0 0 1 623.36 0 235.264 235.264 0 0 1-65.28 461.226666 33.621333 33.621333 0 0 1 0-67.2 168.021333 168.021333 0 0 0 31.573333-333.056 33.578667 33.578667 0 0 1-26.538666-25.685333 257.706667 257.706667 0 0 0-502.869334 0 33.578667 33.578667 0 0 1-26.538666 25.685333 168.064 168.064 0 0 0 31.573333 333.056 33.621333 33.621333 0 0 1 0 67.2 235.178667 235.178667 0 0 1-65.28-461.226666z" data-v-36feed0c></path><path d="M489.386667 522.965333a32 32 0 0 1 45.226666 0l106.666667 106.666667a32 32 0 1 1-45.226667 45.226667l-52.053333-52.010667v264.106667a32 32 0 0 1-64 0v-264.106667l-52.053333 52.053333a32 32 0 1 1-45.226667-45.226666l106.666667-106.666667z" data-v-36feed0c></path></svg><div class="upload-text" data-v-36feed0c> Click or drag to upload an image <div class="upload-hint" data-v-36feed0c> Support PNG, JPG,JPEG. Max 20 MB </div></div></div>',1))]))])],512))}}),he=$(ge,[["__scopeId","data-v-36feed0c"]]),we={class:"main-content"},be={class:"text-center mb-4"},fe={class:"d-flex align-center"},ye=T({__name:"CropImage",setup(B){const E=v("resize");v("zh");const c=v("dimensions"),r=v(""),C=v(""),h=v(!0),m=v("original"),y=v(null),b=v({show:!1,text:"",color:"error"}),f=v(!1),S=o=>{var i;const t=(i=o.target.files)==null?void 0:i[0];if(t){const e=new FileReader;e.onload=n=>{var p;const s=new Image;s.onload=()=>{y.value=s,r.value=s.width.toString(),C.value=s.height.toString()},s.src=(p=n.target)==null?void 0:p.result},e.readAsDataURL(t)}},z=()=>{},u=(o,t="success")=>{b.value={show:!0,text:o,color:t}},x=async o=>{try{if(!y.value)throw new Error("No image loaded");f.value=!0;const t=new FormData,e=await(await fetch(y.value.src)).blob();t.append("image",e);const n={width:o.width,height:o.height,format:o.format==="original"?"":o.format,backgroundColor:o.backgroundColor||"",quality:85,dpi:o.dpi};if(o.targetFileSize){const M={KB:1024,MB:1048576,GB:1073741824},X=o.targetFileSize.size*M[o.targetFileSize.unit];n.targetFileSize=X}t.append("data",JSON.stringify(n));const s=await fetch(`${G}/api/image/resize`,{method:"POST",body:t});if(!s.ok)throw new Error(`HTTP error! status: ${s.status}`);let p="resized_image";const d=s.headers.get("content-disposition");if(d){const M=d.match(/filename="([^"]*)"/);M&&M[1]&&(p=M[1])}const a=await s.blob(),R=window.URL.createObjectURL(a),I=document.createElement("a");I.href=R,I.download=p,document.body.appendChild(I),I.click(),window.URL.revokeObjectURL(R),document.body.removeChild(I),u("图片处理成功！","success")}catch(t){console.error("处理图片失败:",t),u(t instanceof Error?t.message:"处理图片失败，请重试","error")}finally{f.value=!1}},U=o=>{o.ctrlKey&&o.preventDefault()};return W(()=>{window.addEventListener("wheel",U,{passive:!1})}),D(()=>{window.removeEventListener("wheel",U)}),(o,t)=>{const i=k("v-progress-circular"),e=k("v-card"),n=k("v-overlay"),s=k("v-icon"),p=k("v-btn"),d=k("v-snackbar");return V(),Y("main",we,[g(n,{modelValue:f.value,"onUpdate:modelValue":t[0]||(t[0]=a=>f.value=a),class:"align-center justify-center",persistent:"",scrim:"rgba(0,0,0,0.3)"},{default:P(()=>[g(e,{color:"white",width:"300",rounded:"lg",elevation:"8",class:"pa-4"},{default:P(()=>[l("div",be,[g(i,{indeterminate:"",color:"primary",size:"64"})]),t[9]||(t[9]=l("div",{class:"text-center text-body-1"}," 处理图片中... ",-1))]),_:1})]),_:1},8,["modelValue"]),g(ce,{currentTool:E.value,"onUpdate:currentTool":t[1]||(t[1]=a=>E.value=a),resizeMode:c.value,"onUpdate:resizeMode":t[2]||(t[2]=a=>c.value=a),width:r.value,"onUpdate:width":t[3]||(t[3]=a=>r.value=a),height:C.value,"onUpdate:height":t[4]||(t[4]=a=>C.value=a),lockAspectRatio:h.value,"onUpdate:lockAspectRatio":t[5]||(t[5]=a=>h.value=a),exportFormat:m.value,"onUpdate:exportFormat":t[6]||(t[6]=a=>m.value=a),imageElement:y.value,disabled:f.value,onResizeImage:x},null,8,["currentTool","resizeMode","width","height","lockAspectRatio","exportFormat","imageElement","disabled"]),g(he,{imageElement:y.value,onFileUpload:S,onImageClick:z},null,8,["imageElement"]),g(d,{modelValue:b.value.show,"onUpdate:modelValue":t[8]||(t[8]=a=>b.value.show=a),color:b.value.color,timeout:3e3,location:"top",elevation:"4",rounded:"lg",class:"custom-snackbar"},{actions:P(()=>[g(p,{color:"white",variant:"text",onClick:t[7]||(t[7]=a=>b.value.show=!1)},{default:P(()=>t[10]||(t[10]=[_(" 关闭 ")])),_:1})]),default:P(()=>[l("div",fe,[g(s,{icon:b.value.color==="error"?"mdi-alert-circle":"mdi-check-circle",class:"me-2"},null,8,["icon"]),l("span",null,j(b.value.text),1)])]),_:1},8,["modelValue","color"])])}}}),Ce=$(ye,[["__scopeId","data-v-423d145b"]]);export{Ce as C};
