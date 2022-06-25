document.addEventListener('DOMContentLoaded', function(event) {
    var settings = {
        activeSignatures: [],
        burger: document.getElementById('burger'),
        connectionStats: document.getElementById('connection-status'),
        matchesCount: document.getElementById('matches-count').getElementsByTagName('span')[0],
        filtersClear: document.getElementById('filters-clear'),
        filtersCount: document.getElementById('filters-count').getElementsByTagName('span')[0]
    };
    const slugify = (value) => value.toLowerCase().replace(/[^a-z0-9 -]/g, '').replace(/\s+/g, '-').replace(/-+/g, '-');
    const sort = (list) => {      
        signatures = list.getElementsByTagName("li");
        Array.from(signatures)
            .sort((a, b) => parseInt(b.getElementsByClassName('menu-item')[0].getAttribute('data-badge') || 0) - parseInt(a.getElementsByClassName('menu-item')[0].getAttribute('data-badge') || 0))
            .forEach(li => list.appendChild(li));
    };
    const updateStatus = (text, cls) => {
        settings.connectionStats.classList.remove('is-info', 'is-success', 'is-warning', 'is-danger');
        settings.connectionStats.classList.add(cls);
        settings.connectionStats.textContent = text;
    };
    const filterSignature = (signature) => {
        var state = settings.activeSignatures.includes(signature.id);
        signature.classList.toggle('is-active');

        if (!state) settings.activeSignatures.push(signature.id);

        Array.from(document.getElementsByClassName('log')).forEach(log => log.style.display = 'none');
        settings.activeSignatures.forEach((signatureId) => {
            Array.from(document.getElementsByClassName(signatureId)).forEach(log => {
                if (state && signatureId == signature.id) return;

                log.style.display = '';
            });
        });
        
        if (state) {
            settings.activeSignatures.splice(settings.activeSignatures.indexOf(signature.id), 1);
            var anyActive = (settings.activeSignatures.length > 0);
            Array.from(document.getElementsByClassName(anyActive ? signature.id : 'log')).forEach(log => {

                log.style.display = anyActive ? 'none' : '';
            });
        }

        settings.filtersCount.textContent = `${settings.activeSignatures.length} filters`;
    };
    const processEvent = (data) => {
        var eventId = CryptoJS.MD5(JSON.stringify(data)).toString();
        if (document.getElementById(eventId)) return; // duplicate

        var sigId = slugify(data.category);
        switch (true) {
          case sigId.startsWith("cve-"):
            sigId = "cves";
            break;
          case sigId.startsWith("common-web-attack"):
            sigId = "common-web-attacks";
            break;
          default:
             break;
        }

        document.getElementById(sigId) == null ? createSignature(data.category) : true;
        var sigMenuItem = document.getElementById(sigId).getElementsByClassName('menu-item')[0];
        sigMenuItem.setAttribute('data-badge', parseInt(sigMenuItem.getAttribute('data-badge') || 0) + 1);
        sort(document.getElementById('signatures'));

        var row = document.getElementById('messages').insertRow(0);
        row.classList.add('log', sigId);
        row.id = eventId;
        row.insertCell(0).innerHTML = `<td class="date"><span class="datetime" title="${new Date().toLocaleString}">${new Date().toLocaleTimeString()}</span></td>`;
        row.insertCell(1).innerHTML = `<td class="category-name"><strong>${data.category}</strong></td>`;
        row.insertCell(2).innerHTML = `<td class="element"><div>${data.element}</pre></div></td>`;
        row.insertCell(3).innerHTML = `<td class="matches"><strong>${data[data.element]}</strong></td>`;
        row.insertCell(4).innerHTML = `<td class="log-line"><div><pre>${JSON.stringify(data)}</pre></div></td>`;
        settings.matchesCount.textContent = `${document.getElementsByClassName('log').length} threats`;

        if (settings.activeSignatures.length > 0 && !settings.activeSignatures.includes(sigId)) row.style.display = 'none';
    };
    const listenForEvents = () => {
        var endpoint = new URL("http://{{ .TELER_DASHBOARD_HOST }}:{{ .TELER_DASHBOARD_PORT }}");
        endpoint.pathname = "{{ .TELER_DASHBOARD_ENDPOINT }}";
        endpoint.search = "?stream=teler";

        var source = new EventSource(endpoint);

        source.onerror = (e) => {
            updateStatus('Reconnect...', 'is-warning');
        };

        source.onoopen = (e) => {
            updateStatus('Syncing...', 'is-info');
        }

        source.onmessage = (e) => {
            updateStatus('Connected', 'is-success');

            if (document.getElementById('loading')) document.getElementById('loading').remove();
            processEvent(JSON.parse(e.data));
        };
    };
    const createSignature = (name) => {
        var li = document.createElement('li');
        li.id = slugify(name)
        li.innerHTML = `<a href="#" class="menu-item" title="${name}">${name}</a>`;
        li.addEventListener('click', (e) => {
            e.preventDefault();
            filterSignature(li);
        });

        document.getElementById('signatures').appendChild(li);
    };

    (() => {
        burger.addEventListener('click', () => {
            const target = burger.dataset.target;
            const $target = document.getElementById(target);
    
            burger.classList.toggle('is-active');
            $target.classList.toggle('is-active');
        });

        settings.filtersClear.addEventListener('click', (event) => {
            settings.activeSignatures = [];
            settings.filtersCount.textContent = "0 filters";

            Array.from(document.querySelectorAll('#signatures li.is-active')).forEach(log => log.classList.remove('is-active'));
            Array.from(document.getElementsByClassName('log')).forEach(log => log.style.display = '');
        });

        fetch(`signatures.json`)
            .then((resp) => resp.json())
            .then(signatures => {
                signatures.forEach(name => {
                    createSignature(name)
                });

                listenForEvents();
            })
            .catch((err) => {
                alert('Failed to retrieve signatures! Reloading...')
                console.error(err)
            });
    })();
});