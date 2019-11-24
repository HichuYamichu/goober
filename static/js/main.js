function del(resorce) {
  fetch(`/remove/${resorce}`, {
    method: 'delete',
    headers: {
      'x-api-key': 1234
    }
  });
}
