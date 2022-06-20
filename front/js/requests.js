async function send(url, data, on_result) {
    // console.log(url);
    if (data !== undefined && data !== null) {
        response = await fetch(url, {
            method: "POST",
            body: data instanceof FormData ? data : JSON.stringify(data)
        });
    } else {
        response = await fetch(url);
    }
    let status = response.status;
    let message = await response.text();
    try {
        message = JSON.parse(message);
    } catch {}
    // console.log(data);
    // console.log(message);
    if (on_result) await on_result(status, message);
    return [status, message];
};