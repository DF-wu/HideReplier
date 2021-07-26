const botNameInput = document.getElementById("botNameInput");
const colorInput = document.getElementById("colorInput");
const avatarInput = document.getElementById("avatarInput");
const imageInput = document.getElementById("imageInput");
const contentText = document.getElementById("contentText");
const contentDiv = document.querySelector(".content");

const previewThumbnail = document.getElementById("previewThumbnail");
const selectThumbnail = document.getElementById("selectThumbnail");

const preview = document.getElementById("preview");
const submit = document.getElementById("submit");

const ipDisplay = document.getElementById("ipDisplay");
const botNameDisplay = document.getElementById("botNameDisplay");
const infoDisplay = document.getElementById("infoDisplay");
const botNameWithColor = document.getElementById("botNameWithColor");
const contentDisplay = document.getElementById("contentDisplay");
const ipDisplayInContent = document.getElementById("ipDisplayInContent");
const errorDisplay = document.getElementById("errorDisplay");

const ipCountry = document.getElementById("ipCountry");
const ipCity = document.getElementById("ipCity");
const ipISP = document.getElementById("ipISP");

const avatarImg = document.querySelector(".avatar img");
const contentImgDiv = document.querySelector(".contentImgDiv");
const contentImg = document.querySelector(".contentImgDiv img");

const inputs = [botNameInput, colorInput, avatarInput, imageInput, contentText];

const GetIpAPI = "https://httpbin.org/ip";
const GetIpCountry = "https://ipwhois.app/json";
const DefaultAvatarUrl =
    "https://cdn.discordapp.com/avatars/710112845567623238/f377b595ef4e0ea17826d7afbb20633f.webp?size=128";
const BackendUrl = window.location.origin + "/HideBot/discord";

const thumbnails = [
    "https://cdn.discordapp.com/attachments/408969877580414976/714868499994116096/kyaru02.gif",
    "https://cdn.discordapp.com/attachments/591643710454890505/723846090247635084/kyaru01.gif",
    "https://cdn.discordapp.com/attachments/591643710454890505/723846093959594014/kyaru04.gif",
    "https://cdn.discordapp.com/attachments/591643710454890505/723846094353858580/kyaru03.gif",
    "https://cdn.discordapp.com/attachments/591643710454890505/723846098371739728/kyaru08.gif",
    "https://cdn.discordapp.com/attachments/591643710454890505/723846100032946206/peko01.gif",
    "https://cdn.discordapp.com/attachments/591643710454890505/723846103006445618/ue01.gif",
];

let currentIp = null;
let thumbnailLink;

const getIpInfo = async () => {
    const { origin } = await fetch(GetIpAPI).then((res) => res.json());
    const ip = origin.split(",")[0];
    currentIp = ipDisplay.innerHTML = ipDisplayInContent.innerHTML = ip;

    const { country, city, isp } = await fetch(
        GetIpCountry + `/${currentIp}`
    ).then((res) => res.json());

    ipCountry.innerHTML = country;
    ipCity.innerHTML = city;
    ipISP.innerHTML = isp;
};
getIpInfo();

if (localStorage.getItem("color"))
    contentDiv.style.borderColor = colorInput.value =
        localStorage.getItem("color");

const imgOnError = (event) => {
    if (event.target === avatarImg) {
        avatarImg.src = DefaultAvatarUrl;
        errorDisplay.innerHTML = "不支援的頭像連結";
    } else {
        contentImgDiv.hidden = true;
        errorDisplay.innerHTML = "不支援的圖片連結";
    }
};

const previewOption = () => {
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

    //根據選項改預覽小圖
    thumbnailLink = previewThumbnail.src =
        thumbnails[selectThumbnail.selectedIndex];
};

const parseBBCode = (str) => {
    return str.replace("[img]", "").replace("[/img]", "").trim();
};

const submitPost = async () => {
    const postBody = {
        username: botNameInput.value,
        content: contentText.value,
        color: colorInput.value,
        avatar_url: avatarInput.value,
        imgUrl: imageInput.value,
        ip: currentIp,
        thumbnail: thumbnailLink,
    };

    localStorage.setItem("color", colorInput.value);

    await fetch(BackendUrl, {
        body: JSON.stringify(postBody),
        method: "POST",
        mode: "cors",
        headers: {
            "content-type": "application/json",
        },
    });
    alert("發送成功!!");
};

inputs.forEach((inp) =>
    inp.addEventListener("keydown", () => (submit.disabled = true))
);
avatarImg.onerror = imgOnError;
contentImg.onerror = imgOnError;
preview.onclick = previewOption;
submit.onclick = submitPost;
