const postURL = async (url:string) => {
  const rawResponse = await fetch("http://localhost:5000/api/create-hash", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({URL: url}),
  });
  const response = await rawResponse.json();
  return { status: rawResponse.status, body: response };
};

export default postURL