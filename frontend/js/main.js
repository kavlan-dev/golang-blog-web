document.addEventListener('DOMContentLoaded', function() {
    const API_BASE_URL = 'http://localhost:8080/api/posts/';

    // DOM elements
    const postsList = document.getElementById('posts-list');
    const createPostForm = document.getElementById('create-post-form');
    const editPostForm = document.getElementById('edit-post-form');
    const cancelEditBtn = document.getElementById('cancel-edit');

    // Load all posts on page load
    loadPosts();

    // Create post form submission
    createPostForm.addEventListener('submit', function(e) {
        e.preventDefault();

        const title = document.getElementById('post-title').value;
        const content = document.getElementById('post-content').value;

        fetch(API_BASE_URL, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                title: title,
                content: content
            })
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка создания записи');
            }
            return response.json();
        })
        .then(post => {
            // Clear form
            document.getElementById('post-title').value = '';
            document.getElementById('post-content').value = '';

            // Reload posts
            loadPosts();
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Ошибка при создании записи: ' + error.message);
        });
    });

    // Edit post form submission
    editPostForm.addEventListener('submit', function(e) {
        e.preventDefault();

        const id = document.getElementById('edit-post-id').value;
        const title = document.getElementById('edit-post-title').value;
        const content = document.getElementById('edit-post-content').value;

        fetch(API_BASE_URL + id, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                title: title,
                content: content
            })
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка обновления записи');
            }
            return response.json();
        })
        .then(post => {
            // Hide edit form and clear it
            editPostForm.style.display = 'none';
            document.getElementById('edit-post-id').value = '';
            document.getElementById('edit-post-title').value = '';
            document.getElementById('edit-post-content').value = '';

            // Reload posts
            loadPosts();
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Ошибка при обновлении записи: ' + error.message);
        });
    });

    // Cancel edit button
    cancelEditBtn.addEventListener('click', function() {
        editPostForm.style.display = 'none';
        document.getElementById('edit-post-id').value = '';
        document.getElementById('edit-post-title').value = '';
        document.getElementById('edit-post-content').value = '';
    });

    // Load all posts from API
    function loadPosts() {
        fetch(API_BASE_URL)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Ошибка загрузки записей');
                }
                return response.json();
            })
            .then(posts => {
                postsList.innerHTML = '';

                if (posts.length === 0) {
                    postsList.innerHTML = '<div class="alert alert-info">Нет записей в блоге</div>';
                    return;
                }

                posts.forEach(post => {
                    const postElement = document.createElement('div');
                    postElement.className = 'post-item list-group-item';
                    postElement.innerHTML = `
                        <div class="post-title">${post.title}</div>
                        <div class="post-content">${post.content}</div>
                        <div class="post-actions">
                            <button class="btn btn-sm btn-edit" data-id="${post.id}">Редактировать</button>
                            <button class="btn btn-sm btn-delete" data-id="${post.id}">Удалить</button>
                        </div>
                    `;

                    postsList.appendChild(postElement);
                });

                // Add event listeners to edit and delete buttons
                document.querySelectorAll('.btn-edit').forEach(button => {
                    button.addEventListener('click', function() {
                        const postId = this.getAttribute('data-id');
                        loadPostForEdit(postId);
                    });
                });

                document.querySelectorAll('.btn-delete').forEach(button => {
                    button.addEventListener('click', function() {
                        const postId = this.getAttribute('data-id');
                        if (confirm('Вы уверены, что хотите удалить эту запись?')) {
                            deletePost(postId);
                        }
                    });
                });
            })
            .catch(error => {
                console.error('Error:', error);
                postsList.innerHTML = '<div class="alert alert-danger">Ошибка загрузки записей: ' + error.message + '</div>';
            });
    }

    // Load single post for editing
    function loadPostForEdit(postId) {
        fetch(API_BASE_URL + postId)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Ошибка загрузки записи');
                }
                return response.json();
            })
            .then(post => {
                document.getElementById('edit-post-id').value = post.id;
                document.getElementById('edit-post-title').value = post.title;
                document.getElementById('edit-post-content').value = post.content;
                editPostForm.style.display = 'block';
            })
            .catch(error => {
                console.error('Error:', error);
                alert('Ошибка при загрузке записи для редактирования: ' + error.message);
            });
    }

    // Delete post
    function deletePost(postId) {
        fetch(API_BASE_URL + postId, {
            method: 'DELETE'
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка удаления записи');
            }
            // Reload posts
            loadPosts();
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Ошибка при удалении записи: ' + error.message);
        });
    }
});
