/**
 * UI взаимодействие и обработка событий
 */
document.addEventListener('DOMContentLoaded', function() {
    // Инициализация переключателя темы
    initThemeToggle();

    // Инициализация обработчиков событий
    initEventHandlers();

    // Загрузка постов при загрузке страницы
    loadPosts();
});

/**
 * Инициализация переключателя темы
 */
function initThemeToggle() {
    const themeToggle = document.getElementById('theme-toggle');
    const themeStylesheet = document.getElementById('theme-stylesheet');

    // Проверяем сохраненные настройки темы
    const savedTheme = localStorage.getItem('blogTheme');
    if (savedTheme === 'dark') {
        themeStylesheet.href = 'css/dark-theme.css';
        themeToggle.innerHTML = '<i class="bi bi-sun-fill"></i> Светлая тема';
    }

    themeToggle.addEventListener('click', function() {
        if (themeStylesheet.href.includes('dark-theme.css')) {
            themeStylesheet.href = 'css/styles.css';
            themeToggle.innerHTML = '<i class="bi bi-moon-fill"></i> Темная тема';
            localStorage.setItem('blogTheme', 'light');
        } else {
            themeStylesheet.href = 'css/dark-theme.css';
            themeToggle.innerHTML = '<i class="bi bi-sun-fill"></i> Светлая тема';
            localStorage.setItem('blogTheme', 'dark');
        }
    });
}

/**
 * Инициализация обработчиков событий
 */
function initEventHandlers() {
    // Создание поста
    document.getElementById('savePostBtn').addEventListener('click', handleCreatePost);

    // Обновление поста
    document.getElementById('updatePostBtn').addEventListener('click', handleUpdatePost);

    // Подтверждение удаления поста
    document.getElementById('confirmDeleteBtn').addEventListener('click', handleDeletePost);
}

/**
 * Загрузка и отображение постов
 */
async function loadPosts() {
    const postsContainer = document.getElementById('posts-container');

    try {
        const posts = await blogApi.getAllPosts();

        if (posts.length === 0) {
            postsContainer.innerHTML = `
                <div class="col-12 text-center">
                    <div class="alert alert-info">
                        Нет записей. Создайте первую запись!
                    </div>
                </div>
            `;
            return;
        }

        // Сортировка постов по дате создания (новые сверху)
        posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));

        postsContainer.innerHTML = posts.map(post => `
            <div class="col-md-6 col-lg-4">
                <div class="card post-card">
                    <div class="card-header">
                        <h5 class="card-title mb-0">${post.title}</h5>
                    </div>
                    <div class="card-body">
                        <p class="card-text">${post.content}</p>
                        <div class="post-meta">
                            Создано: ${blogApi.formatDate(post.created_at)}<br>
                            Обновлено: ${blogApi.formatDate(post.updated_at)}
                        </div>
                    </div>
                    <div class="card-footer">
                        <div class="post-actions">
                            <button class="btn btn-sm btn-outline-primary edit-post-btn" data-post-id="${post.id}">
                                <i class="bi bi-pencil"></i> Редактировать
                            </button>
                            <button class="btn btn-sm btn-outline-danger delete-post-btn" data-post-id="${post.id}">
                                <i class="bi bi-trash"></i> Удалить
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        `).join('');

        // Инициализация обработчиков для кнопок редактирования и удаления
        initPostActionButtons();

    } catch (error) {
        postsContainer.innerHTML = `
            <div class="col-12 text-center">
                <div class="alert alert-danger">
                    Ошибка загрузки постов: ${error.message}
                </div>
            </div>
        `;
    }
}

/**
 * Инициализация обработчиков для кнопок редактирования и удаления
 */
function initPostActionButtons() {
    // Кнопки редактирования
    document.querySelectorAll('.edit-post-btn').forEach(button => {
        button.addEventListener('click', function() {
            const postId = this.getAttribute('data-post-id');
            showEditPostModal(postId);
        });
    });

    // Кнопки удаления
    document.querySelectorAll('.delete-post-btn').forEach(button => {
        button.addEventListener('click', function() {
            const postId = this.getAttribute('data-post-id');
            showDeletePostModal(postId);
        });
    });
}

/**
 * Показать модальное окно редактирования поста
 * @param {number} postId ID поста
 */
async function showEditPostModal(postId) {
    try {
        const post = await blogApi.getPostById(postId);

        document.getElementById('editPostId').value = post.id;
        document.getElementById('editPostTitle').value = post.title;
        document.getElementById('editPostContent').value = post.content;

        const editModal = new bootstrap.Modal(document.getElementById('editPostModal'));
        editModal.show();
    } catch (error) {
        blogApi.showNotification(`Ошибка загрузки поста: ${error.message}`, 'error');
    }
}

/**
 * Показать модальное окно подтверждения удаления
 * @param {number} postId ID поста
 */
function showDeletePostModal(postId) {
    const deleteBtn = document.getElementById('confirmDeleteBtn');
    deleteBtn.setAttribute('data-post-id', postId);

    const deleteModal = new bootstrap.Modal(document.getElementById('deletePostModal'));
    deleteModal.show();
}

/**
 * Обработка создания поста
 */
async function handleCreatePost() {
    const title = document.getElementById('postTitle').value.trim();
    const content = document.getElementById('postContent').value.trim();

    if (!title || !content) {
        blogApi.showNotification('Пожалуйста, заполните все поля', 'error');
        return;
    }

    try {
        await blogApi.createPost(title, content);
        blogApi.showNotification('Запись успешно создана!', 'success');

        // Закрытие модального окна
        const createModal = bootstrap.Modal.getInstance(document.getElementById('createPostModal'));
        createModal.hide();

        // Очистка формы
        document.getElementById('createPostForm').reset();

        // Обновление списка постов
        loadPosts();
    } catch (error) {
        blogApi.showNotification(`Ошибка создания записи: ${error.message}`, 'error');
    }
}

/**
 * Обработка обновления поста
 */
async function handleUpdatePost() {
    const postId = document.getElementById('editPostId').value;
    const title = document.getElementById('editPostTitle').value.trim();
    const content = document.getElementById('editPostContent').value.trim();

    if (!title || !content) {
        blogApi.showNotification('Пожалуйста, заполните все поля', 'error');
        return;
    }

    try {
        await blogApi.updatePost(postId, title, content);
        blogApi.showNotification('Запись успешно обновлена!', 'success');

        // Закрытие модального окна
        const editModal = bootstrap.Modal.getInstance(document.getElementById('editPostModal'));
        editModal.hide();

        // Обновление списка постов
        loadPosts();
    } catch (error) {
        blogApi.showNotification(`Ошибка обновления записи: ${error.message}`, 'error');
    }
}

/**
 * Обработка удаления поста
 */
async function handleDeletePost() {
    const postId = document.getElementById('confirmDeleteBtn').getAttribute('data-post-id');

    try {
        await blogApi.deletePost(postId);
        blogApi.showNotification('Запись успешно удалена!', 'success');

        // Закрытие модального окна
        const deleteModal = bootstrap.Modal.getInstance(document.getElementById('deletePostModal'));
        deleteModal.hide();

        // Обновление списка постов
        loadPosts();
    } catch (error) {
        blogApi.showNotification(`Ошибка удаления записи: ${error.message}`, 'error');
    }
}

// Экспорт функций для использования в других модулях
window.blogUi = {
    loadPosts,
    showEditPostModal,
    showDeletePostModal,
    handleCreatePost,
    handleUpdatePost,
    handleDeletePost
};
