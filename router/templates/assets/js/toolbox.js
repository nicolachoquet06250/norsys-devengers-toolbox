function match(regex, str) {
    let m;

    const matches = [];
    let i = 0;

     while ((m = regex.exec(str)) !== null) {
        if (m.index === regex.lastIndex) {
            regex.lastIndex++;
        }
    
        const match = [];

        m.forEach((_match, groupIndex) => {
            console.log(`Found match, group ${groupIndex}: ${_match}`);
            match[groupIndex] = _match;
        });

        matches.push(match);
     }

     return matches;
}

function createToolEl(tool) {
    /**
     * @type {HTMLElement}
     */
    const toolEl = document.querySelector('#tool').content.cloneNode(true);
    const regex = /\{([a-z]+)\}/g;

    const id = tool[toolEl.querySelector('div').getAttribute(':id')];
    toolEl.querySelector('div').id = id.toLowerCase().replace(/[ '".]/g, '-');
    toolEl.querySelector('div').removeAttribute(':id');

    const target = tool[toolEl.querySelector('.tool').getAttribute(':target')];
    toolEl.querySelector('.tool').setAttribute('href', target);
    toolEl.querySelector('.tool').setAttribute('target', '_blank');
    toolEl.querySelector('.tool').dataset.href = target;
    toolEl.querySelector('.tool').removeAttribute(':target');
    toolEl.querySelector('.tool').addEventListener('click', e => {
        if (e.target.tagName === 'IMG' || e.target.tagName === 'BUTTON') {
            e.preventDefault();
            e.stopPropagation();
            e.stopImmediatePropagation();
            return;
        }
    })

    const oldImageSrc = toolEl.querySelector('div > img').getAttribute('src');
    let _match = match(regex, oldImageSrc).length === 1 
        ? match(regex, oldImageSrc)[0] : match(regex, oldImageSrc);
    const imageSrc = oldImageSrc.replace(
        _match[0], tool[_match[1]]
    );

    const oldImageAlt = toolEl.querySelector('div > img').getAttribute('alt');
    _match = match(regex, oldImageAlt).length === 1 
        ? match(regex, oldImageAlt)[0] : match(regex, oldImageAlt);
    const imageAlt = oldImageAlt.replace(
        _match[0], tool[_match[1]]
    );

    const oldH2Content = toolEl.querySelector('h2').textContent ?? '';
    _match = match(regex, oldH2Content).length === 1 
        ? match(regex, oldH2Content)[0] : match(regex, oldH2Content);
    const h2Content = oldH2Content.replace(
        _match[0], tool[_match[1]]
    );

    const h3Cond = !!tool[toolEl.querySelector('h3').getAttribute(':if')];
    const oldH3Content = toolEl.querySelector('h3').textContent ?? '';
    _match = match(regex, oldH3Content).length === 1 
        ? match(regex, oldH3Content)[0] : match(regex, oldH3Content);
    const h3Content = oldH3Content.replace(
        _match[0], tool[_match[1]]
    );

    const descriptionCond = !!tool[toolEl.querySelector('.description').getAttribute(':if')];
    const oldDescriptionContent = toolEl.querySelector('.description').textContent ?? '';
    _match = match(regex, oldDescriptionContent).length === 1 
        ? match(regex, oldDescriptionContent)[0] : match(regex, oldDescriptionContent);
    const descriptionContent = oldDescriptionContent.replace(
        _match[0], tool[_match[1]]
    );

    const menuLinks = toolEl.querySelectorAll('nav > ul > li > a');
    [...menuLinks].map(link => {
        const oldLinkContent = link.textContent ?? '';
        
        let linkContent = oldLinkContent;
        let matches = match(regex, linkContent);
        for (const match of matches) {
            linkContent = linkContent.replace(
                match[0], 
                tool[match[1]]
            );
        }

        link.textContent = linkContent;

        const oldLinkHref = link.getAttribute(':target');
        let linkHref = oldLinkHref;
        matches = match(regex, linkHref);
        for (const match of matches) {
            if (_match[0]) {
                linkHref = linkHref.replace(
                    match[0], tool[match[1]]
                );
            }
        }
        link.removeAttribute(':target');
        link.setAttribute('href', linkHref);
    });

    toolEl.querySelector('div > img').setAttribute('src', imageSrc);
    toolEl.querySelector('div > img').setAttribute('alt', imageAlt);
    toolEl.querySelector('h2').textContent = h2Content;
    if (h3Cond) {
        toolEl.querySelector('h3').removeAttribute(':if');
        toolEl.querySelector('h3').textContent = h3Content;
    } else {
        toolEl.querySelector('h3').remove();
    }
    if (descriptionCond) {
        toolEl.querySelector('.description').removeAttribute(':if');
        toolEl.querySelector('.description').textContent = descriptionContent;
    } else {
        toolEl.querySelector('.description').remove();
    }

    toolEl.querySelector('button.like').addEventListener('click', (e) => {
        const img = e.currentTarget.querySelector('img');
        if (img.dataset.icon_state === 'like') {
            img.dataset.icon_state = 'unlike';
            img.setAttribute('src', 'https://fonts.gstatic.com/s/i/short-term/release/materialsymbolsoutlined/favorite/fill1/48px.svg');
        } else {
            img.dataset.icon_state = 'like';
            img.setAttribute('src', 'https://fonts.gstatic.com/s/i/short-term/release/materialsymbolsoutlined/favorite/default/48px.svg');
        }
    });

    const diablog = toolEl.querySelector('nav');
    toolEl.querySelector('button.menu').addEventListener('click', () => {
        diablog.classList.toggle('show');
    });

    return toolEl;
}

window.addEventListener('load', () => {
    fetch('/api/tools').then(r => r.json()).then(json => {
        for (const tool of json) {
            document.querySelector('.container').append(createToolEl(tool));
        }
    }).catch(err => {
        const error = document.querySelector('.error');
        error.innerHTML = err.message;
        error.classList.add('show');

        setTimeout(() => {
            error.classList.remove('show');
        }, 5000);
    });
});

document.addEventListener('click', () => {
    [...toolEl.querySelectorAll('.container > nav')].map(dialog => {
        dialog.classList.remove('show');
    })
});

function writeInClipboard(name, url) {
    sendMessage({
        channel: 'Clipboard',
        data: {
            text: url
        }
    });

    const success = document.querySelector('.success');
    success.innerHTML = `L'url de ${name} à été copié.`;
    success.classList.add('show');
    document.querySelector(`#${name.toLowerCase().replace(/[ '".]/g, '-')} nav`).classList.remove('show');

    setTimeout(() => {
        success.classList.remove('show');
    }, 5000);
}
