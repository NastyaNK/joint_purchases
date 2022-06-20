class ImageSlider {
    constructor(slider, image, template_) {
        this.slider = slider;
        this.image = image;
        // console.log(this.slider, this.image)
        this.template_ = template_;
        this.first_img = true;
    }
    async add_image(image, template_ = null) {
        let div_ = document.createElement("div");
        div_.innerHTML = await template(template_ == null ? this.template_ : template_, { "image": image });
        div_ = div_.children[0];
        div_.onclick = () => {
            this.onclick(div_);
        }
        this.slider.appendChild(div_);
        // this.slider.innerHTML += img_block.slice(0, i) + " onclick=\""+this.pname+".onclick(this);\""+img_block.slice(i);
        if (this.first_img) {
            this.onclick(this.slider.children[0]);
            this.first_img = false;
        }
        // console.log(this.slider, this.image)
    }
    onclick(block) {
        let image = block.querySelector('img[forslider]');
        this.image.src = image.src;
        let blocks = this.slider.children;
        for (let i = 0; i < blocks.length; i++)
            blocks[i].classList.remove("slider-image-active");
        block.classList.add("slider-image-active");
    }
}


function centering(obj, target, offset = 0) {
    obj.style.left = offset + target.offsetLeft + 'px';
    obj.style.top = target.offsetTop + target.clientHeight / 2 - obj.clientHeight / 2 + 'px';
}