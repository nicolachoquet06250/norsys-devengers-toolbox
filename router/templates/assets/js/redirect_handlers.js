document.addEventListener('astilectron-ready', () => {
    /**
             * @param {string} message.channel
             * @param {Record<string, string>|Record<string, string>[]} message.data
             * @returns {string}
             */
    const handleNewMessage = message => {
        switch (message.channel) {
            case 'Redirect':
                window.location.href = message.data.uri;
                return message.data.uri;
        }
    };
    astilectron.onMessage(handleNewMessage);
});