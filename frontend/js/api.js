/**
 * API взаимодействие с бэкендом блога
 */
const API_BASE_URL = 'http://localhost:8080/api/posts/';

/**
 * Получение всех постов
 * @returns {Promise<Array>} Список постов
 */
async function getAllPosts() {
    try {
        const response = await fetch(API_BASE_URL);
        if (!response.ok) {
            throw new Error(`Ошибка загрузки постов: ${response.status}`);
        }
        return await response.json();
    } catch (error) {
        console.error('Ошибка при получении постов:', error);
        throw error;
    }
}

/**
 * Получение одного поста по ID
 * @param {number} postId ID поста
 * @returns {Promise<Object>} Объект поста
 */
async function getPostById(postId) {
    try {
        const response = await fetch(`${API_BASE_URL}${postId}`);
        if (!response.ok) {
            throw new Error(`Ошибка получения поста: ${response.status}`);
        }
        return await response.json();
    } catch (error) {
        console.error(`Ошибка при получении поста ${postId}:`, error);
        throw error;
    }
}

/**
 * Создание нового поста
 * @param {string} title Заголовок поста
 * @param {string} content Содержимое поста
 * @returns {Promise<Object>} Созданный пост
 */
async function createPost(title, content) {
    try {
        const response = await fetch(API_BASE_URL, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                title: title,
                content: content
            })
        });

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({}));
            throw new Error(errorData.message || `Ошибка создания поста: ${response.status}`);
        }

        return await response.json();
    } catch (error) {
        console.error('Ошибка при создании поста:', error);
        throw error;
    }
}

/**
 * Обновление поста
 * @param {number} postId ID поста
 * @param {string} title Заголовок поста
 * @param {string} content Содержимое поста
 * @returns {Promise<Object>} Обновленный пост
 */
async function updatePost(postId, title, content) {
    try {
        const response = await fetch(`${API_BASE_URL}${postId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                title: title,
                content: content
            })
        });

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({}));
            throw new Error(errorData.message || `Ошибка обновления поста: ${response.status}`);
        }

        return await response.json();
    } catch (error) {
        console.error(`Ошибка при обновлении поста ${postId}:`, error);
        throw error;
    }
}

/**
 * Удаление поста
 * @param {number} postId ID поста
 * @returns {Promise<void>}
 */
async function deletePost(postId) {
    try {
        const response = await fetch(`${API_BASE_URL}${postId}`, {
            method: 'DELETE'
        });

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({}));
            throw new Error(errorData.message || `Ошибка удаления поста: ${response.status}`);
        }
    } catch (error) {
        console.error(`Ошибка при удалении поста ${postId}:`, error);
        throw error;
    }
}

/**
 * Форматирование даты
 * @param {string} dateString Строка даты
 * @returns {string} Отформатированная дата
 */
function formatDate(dateString) {
    const date = new Date(dateString);
    return date.toLocaleString('ru-RU', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
    });
}

/**
 * Показать уведомление
 * @param {string} message Сообщение
 * @param {string} type Тип (success, error, info)
 */
function showNotification(message, type = 'info') {
    const notification = document.createElement('div');
    notification.className = `alert alert-${type} alert-dismissible fade show`;
    notification.role = 'alert';
    notification.innerHTML = `
        ${message}
        <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Закрыть"></button>
    `;

    const container = document.querySelector('.container');
    container.prepend(notification);

    setTimeout(() => {
        notification.remove();
    }, 5000);
}

// Экспорт функций для использования в других модулях
window.blogApi = {
    getAllPosts,
    getPostById,
    createPost,
    updatePost,
    deletePost,
    formatDate,
    showNotification
};
