function del(resorce) {
  fetch(`/remove/${resorce}`, {
    method: 'delete'
  });
}
