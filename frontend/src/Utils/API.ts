
let apiUrl = "http://localhost:8080/"

// ***************************
// ОСТОРОЖНО, файл совсем без типизации :(
// TODO здесь типы еще СОВСЕМ ВООБЩЕ не расписал, все время ушло на отладку бэка и фронтА
// ***************************

export const apiGetRequest = async ({
	path,
	method = 'GET',
	fetch = window.fetch,
}) => {

	let request, response, apiError;
	try {
		request = await fetch(apiUrl + path, {
			method,
		});

		response = await request.text();
		// response = await request.json();
	} catch (error) {
		console.log('CATCH ERROR', error);
		apiError = error;
	}

	if (!request) {
		return { request: { status: 500 }, response: {}, apiError: apiError };
	}

	console.log('STATUS of apiGetRequest:', request.status);
	// console.log('RESPONSE of apiGetRequest:', response);

	apiError = handleApiError(request.status, response);

	if (apiError !== '') {
		console.log('apiError from apiGetRequest.js: ', apiError);
	}

	return { request, response, apiError };
};

export const apiPutRequest = async ({
	path,
	data,
	json = true,
	method = 'PUT',
	fetch = window.fetch,
}) => {
	const headers = {};
	if (json) {
		data = JSON.stringify(data);
		headers['Content-Type'] = 'application/json';
		// headers['Content-Type'] = 'text/html; charset=ascii';
	}

	let request, response, apiError;
	try {
		request = await fetch(apiUrl + path, {
			method,
			headers,
			body: data,
		});

		// response = await request.text();
		response = await request.json();
	} catch (error) {
		console.log('CATCH ERROR', error);
		apiError = error;
	}

	if (!request) {
		return { request: { status: 500 }, response: {}, apiError: apiError };
	}

	console.log('STATUS of apiPutRequest:', request.status);
	// console.log('RESPONSE of apiRequest GET:', response);

	apiError = handleApiError(request.status, response);

	if (apiError !== '') {
		console.log('apiError from apiPutRequest.js: ', apiError);
	}

	return { request, response, apiError };
};

export const apiDeleteRequest = async ({
	path,
	data,
	json = true,
	method = 'DELETE',
	fetch = window.fetch,
}) => {
	const headers = {};
	if (json) {
		data = JSON.stringify(data);
		headers['Content-Type'] = 'application/json';
	}

	let request, response, apiError;
	try {
		request = await fetch(apiUrl + path, {
			method,
			headers,
			body: data,
		});

		response = await request.text();
		// response = await request.json();
	} catch (error) {
		console.log('CATCH ERROR', error);
		apiError = error;
	}

	if (!request) {
		return { request: { status: 500 }, response: {}, apiError: apiError };
	}

	console.log('STATUS of apiDeleteRequest:', request.status);

	apiError = handleApiError(request.status, response);

	if (apiError !== '') {
		console.log('apiError from apiDeleteRequest.js: ', apiError);
	}

	return { request, response, apiError };
};

export const handleApiError = (status, response) => {
	let error = '';
	if (status === 401) {
		error = response || 'Login required to proceed.';
	} else if (status !== 200 && status !== 201) {
		error = response;
	}
	// console.log(status, error);

	return error;
};
