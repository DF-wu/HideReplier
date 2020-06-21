const botNameInput = document.getElementById('botNameInput');
const colorInput = document.getElementById('colorInput');
const avatarInput = document.getElementById('avatarInput');
const imageInput = document.getElementById('imageInput');
const contentText = document.getElementById('contentText');
const contentDiv = document.querySelector('.content');

const preview = document.getElementById('preview');
const submit = document.getElementById('submit');

const ipDisplay = document.getElementById('ipDisplay');
const botNameDisplay = document.getElementById('botNameDisplay');
const infoDisplay = document.getElementById('infoDisplay');
const botNameWithColor = document.getElementById('botNameWithColor');
const contentDisplay = document.getElementById('contentDisplay');
const ipDisplayInContent = document.getElementById('ipDisplayInContent');
const errorDisplay = document.getElementById('errorDisplay');

const ipCountry = document.getElementById('ipCountry');
const ipCity = document.getElementById('ipCity');
const ipISP = document.getElementById('ipISP');

const avatarImg = document.querySelector('.avatar img');
const contentImgDiv = document.querySelector('.contentImgDiv');
const contentImg = document.querySelector('.contentImgDiv img');


const inputs = [
    botNameInput,
    colorInput,
    avatarInput,
    imageInput,
    contentText
];

const CorsAnyWhere = "https://cors-anywhere.herokuapp.com/";
const GetIpAPI = "https://httpbin.org/ip";
const GetIpCountry = "http://ip-api.com/json"
const DefaultAvatarUrl = "https://cdn.discordapp.com/avatars/710112845567623238/f377b595ef4e0ea17826d7afbb20633f.webp?size=128";
const BackendUrl = window.location.origin + "/HideBot/discord";
let currentIp = null;


fetch(GetIpAPI).then(res => res.json()).then(res => {
    const ip = res.origin.split(',')[0];
    currentIp = ipDisplay.innerHTML = ipDisplayInContent.innerHTML = ip;
    

    fetch(GetIpCountry + `/${currentIp}`).then(res => res.json()).then(res => {
        ipCountry.innerHTML = res.country;
        ipCity.innerHTML = res.city;
        ipISP.innerHTML = res.isp;
    });
});



if (localStorage.getItem("color")) contentDiv.style.borderColor = colorInput.value = localStorage.getItem("color");


function imgOnError(event) {
    if (event.target === avatarImg) {
        avatarImg.src = DefaultAvatarUrl;
        errorDisplay.innerHTML = "不支援的頭像連結";
    } else {
        contentImgDiv.hidden = true;
        errorDisplay.innerHTML = "不支援的圖片連結";
    }
}

function previewOption() {
    errorDisplay.innerHTML = "";
    if (!botNameInput.value || !contentText.value) {
        errorDisplay.innerHTML = "機器人名字以及內容不可為空";
        return;
    }

    contentImgDiv.hidden = true;
    botNameDisplay.innerHTML = botNameWithColor.innerHTML = botNameInput.value;
    contentDiv.style.borderColor = colorInput.value;
    contentDisplay.innerHTML = contentText.value;

    const avatarUrl = parseBBCode(avatarInput.value.trim());
    const imgUrl = parseBBCode(imageInput.value.trim());

    if (avatarUrl) avatarImg.src = avatarUrl;
    if (imgUrl) {
        contentImgDiv.hidden = false;
        contentImg.src = imgUrl;
    }

    if (!errorDisplay.innerHTML) {
        submit.disabled = false;
    }
}

function parseBBCode(str) {
    return str.replace("[img]","").replace("[/img]","").trim();
}

function submitPost() {
    const postBody = {
        username: botNameInput.value,
        content: contentText.value,
        color: colorInput.value,
        avatar_url: avatarInput.value,
        imgUrl: imageInput.value,
        ip: currentIp,
    };
    localStorage.setItem("color", colorInput.value);

    fetch(BackendUrl, {
        body: JSON.stringify(postBody),
        method: 'POST',
        mode: 'cors',
        headers: {
            "content-type": "application/json"
        }
    }).then(res => res.json()).then(res => alert("發送成功!!"));
}

inputs.forEach(inp => inp.addEventListener('keydown', () => submit.disabled = true));
avatarImg.onerror = imgOnError;
contentImg.onerror = imgOnError;
preview.addEventListener('click', previewOption);
submit.addEventListener('click', submitPost);