const botNameInput = document.getElementById("botNameInput");
const colorInput = document.getElementById("colorInput");
const avatarInput = document.getElementById("avatarInput");
const imageInput = document.getElementById("imageInput");
const contentText = document.getElementById("contentText");
const contentDiv = document.querySelector(".content");
const infoDisplay = document.getElementById("infoDisplay");

const selectThumbnail = document.getElementById("selectThumbnail");
const thumbnailPicker = document.getElementById("thumbnailPicker");

const preview = document.getElementById("preview");
const submit = document.getElementById("submit");

const ipDisplay = document.getElementById("ipDisplay");
const botNameDisplay = document.getElementById("botNameDisplay");
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
const previewThumbnailImg = document.getElementById("previewThumbnail");

const inputs = [botNameInput, colorInput, avatarInput, imageInput, contentText];

const GetIpAPI = "https://httpbin.org/ip";
const GetIpCountry = "https://ipwhois.app/json";
const DefaultAvatarUrl =
    "https://cdn.discordapp.com/avatars/710112845567623238/f377b595ef4e0ea17826d7afbb20633f.webp?size=128";
const BackendUrl = `${window.location.origin}/HideBot/discord`;
const ThumbnailBaseUrl = `${window.location.origin}/thumbs`;
const VersionUrl = `${window.location.origin}/HideBot/discord/version`;
const thumbnails = [
    `${ThumbnailBaseUrl}/01.gif`,
    `${ThumbnailBaseUrl}/02.gif`,
    `${ThumbnailBaseUrl}/03.gif`,
    `${ThumbnailBaseUrl}/04.gif`,
    `${ThumbnailBaseUrl}/05.gif`,
    `${ThumbnailBaseUrl}/06.gif`,
    `${ThumbnailBaseUrl}/07.gif`
];

let currentIp = null;
let thumbnailLink = thumbnails[0];
let submitting = false;
let hasPendingChanges = true;

const setText = (element, value) => {
    element.textContent = value;
};

const setStatus = (message, state = "neutral") => {
    setText(errorDisplay, message);
    errorDisplay.dataset.state = state;
};

const markPreviewDirty = () => {
    if (submitting) {
        return;
    }

    hasPendingChanges = true;
    submit.disabled = true;
    if (errorDisplay.dataset.state === "success") {
        setStatus("", "neutral");
    }
};

const setIpFallback = () => {
    setText(ipDisplay, currentIp || "保護中");
    setText(ipDisplayInContent, currentIp || "保護中");
    setText(ipCountry, "定位保護中");
    setText(ipCity, "定位保護中");
    setText(ipISP, "定位保護中");
};

const isValidHttpUrl = (value) => {
    if (!value) {
        return true;
    }

    try {
        const parsed = new URL(value);
        return parsed.protocol === "http:" || parsed.protocol === "https:";
    } catch (_error) {
        return false;
    }
};

const setThumbnailSelection = (index) => {
    selectThumbnail.selectedIndex = index;
    thumbnailLink = thumbnails[index];
    previewThumbnailImg.src = thumbnailLink;
    document.querySelectorAll(".thumbnailOption").forEach((button, buttonIndex) => {
        button.classList.toggle("is-active", buttonIndex === index);
        button.setAttribute("aria-pressed", buttonIndex === index ? "true" : "false");
    });
};

const buildThumbnailPicker = () => {
    thumbnailPicker.innerHTML = "";
    thumbnails.forEach((thumbnail, index) => {
        const button = document.createElement("button");
        button.type = "button";
        button.className = "thumbnailOption";
        button.setAttribute("aria-label", `預設小圖 ${index + 1}`);

        const image = document.createElement("img");
        image.src = thumbnail;
        image.alt = `預設小圖 ${index + 1}`;

        const label = document.createElement("span");
        label.className = "thumbnailOptionLabel";
        label.textContent = `#${index + 1}`;

        button.append(image, label);
        button.addEventListener("click", () => {
            setThumbnailSelection(index);
            markPreviewDirty();
        });
        thumbnailPicker.append(button);
    });
};

const loadVersion = async () => {
    try {
        const version = await fetch(VersionUrl).then((res) => res.json());
        setText(infoDisplay, `匿名機器人v${version}（點我去發文）`);
    } catch (_error) {
        setText(infoDisplay, "匿名機器人（版本讀取失敗）");
    }
};

const getIpInfo = async () => {
    try {
        const { origin } = await fetch(GetIpAPI).then((res) => res.json());
        const ip = origin.split(",")[0];
        currentIp = ip;
        setText(ipDisplay, ip);
        setText(ipDisplayInContent, ip);

        const { country, city, isp } = await fetch(`${GetIpCountry}/${currentIp}`).then(
            (res) => res.json()
        );

        setText(ipCountry, country || "定位保護中");
        setText(ipCity, city || "定位保護中");
        setText(ipISP, isp || "定位保護中");
    } catch (_error) {
        setIpFallback();
    }
};
getIpInfo();
buildThumbnailPicker();
loadVersion();

const savedColor = localStorage.getItem("color") || "#7c5cff";
contentDiv.style.borderColor = colorInput.value = savedColor;
setThumbnailSelection(0);
contentImgDiv.hidden = true;

const imgOnError = (event) => {
    if (event.target === avatarImg) {
        avatarImg.src = DefaultAvatarUrl;
        setStatus("不支援的頭像連結", "error");
    } else {
        contentImgDiv.hidden = true;
        setStatus("不支援的圖片連結", "error");
    }

    submit.disabled = true;
};

const previewOption = () => {
    setStatus("", "neutral");
    if (!botNameInput.value || !contentText.value) {
        setStatus("機器人名字以及內容不可為空", "error");
        submit.disabled = true;
        return;
    }

    contentImgDiv.hidden = true;
    setText(botNameDisplay, botNameInput.value);
    setText(botNameWithColor, botNameInput.value);
    contentDiv.style.borderColor = colorInput.value;
    setText(contentDisplay, contentText.value);

    const avatarUrl = parseBBCode(avatarInput.value.trim());
    const imgUrl = parseBBCode(imageInput.value.trim());

    if (!isValidHttpUrl(avatarUrl)) {
        setStatus("頭像連結格式不正確", "error");
        submit.disabled = true;
        return;
    }

    if (!isValidHttpUrl(imgUrl)) {
        setStatus("圖片連結格式不正確", "error");
        submit.disabled = true;
        return;
    }

    avatarImg.src = avatarUrl || DefaultAvatarUrl;
    if (imgUrl) {
        contentImgDiv.hidden = false;
        contentImg.src = imgUrl;
    }

    if (!errorDisplay.textContent && hasPendingChanges) {
        submit.disabled = false;
    }
};

const parseBBCode = (str) => {
    return str.replace("[img]", "").replace("[/img]", "").trim();
};

const submitPost = async () => {
    submitting = true;
    submit.disabled = true;
    const originalLabel = submit.textContent;
    submit.textContent = "發送中...";
    const postBody = {
        username: botNameInput.value,
        content: contentText.value,
        color: colorInput.value,
        avatar_url: parseBBCode(avatarInput.value.trim()),
        imgUrl: parseBBCode(imageInput.value.trim()),
        ip: currentIp,
        thumbnail: thumbnailLink,
    };

    localStorage.setItem("color", colorInput.value);

    try {
        const response = await fetch(BackendUrl, {
            body: JSON.stringify(postBody),
            method: "POST",
            mode: "cors",
            headers: {
                "content-type": "application/json",
            },
        });

        if (!response.ok) {
            setStatus("發送失敗，請稍後再試一次", "error");
            return;
        }

        setStatus("發送成功，訊息已送出。", "success");
        hasPendingChanges = false;
        submit.disabled = true;
    } catch (_error) {
        setStatus("連線失敗，請稍後再試一次", "error");
        submit.disabled = false;
    } finally {
        submitting = false;
        submit.textContent = originalLabel;
        if (errorDisplay.dataset.state !== "success") {
            submit.disabled = false;
        }
    }
};

inputs.forEach((inp) => {
    inp.addEventListener("input", markPreviewDirty);
    inp.addEventListener("change", markPreviewDirty);
});
selectThumbnail.addEventListener("change", () => {
    setThumbnailSelection(selectThumbnail.selectedIndex);
    markPreviewDirty();
});
previewThumbnailImg.onerror = () => {
    setThumbnailSelection(0);
};
avatarImg.onerror = imgOnError;
contentImg.onerror = imgOnError;
preview.onclick = previewOption;
submit.onclick = submitPost;
