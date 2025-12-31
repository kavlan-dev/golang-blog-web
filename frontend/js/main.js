/**
 * Основной файл приложения
 * Инициализация и настройка фронтенда
 */

// Импорт Bootstrap иконок
document.addEventListener('DOMContentLoaded', function () {
    // Динамическая загрузка Bootstrap иконок
    loadBootstrapIcons();

    // Инициализация приложения
    initApp();
});

/**
 * Загрузка Bootstrap иконок
 */
function loadBootstrapIcons() {
    const link = document.createElement('link');
    link.rel = 'stylesheet';
    link.href = 'https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.0/font/bootstrap-icons.css';
    document.head.appendChild(link);
}

/**
 * Инициализация приложения
 */
function initApp() {
    console.log('Блог приложение запущено!');

    // Настройка обработки ошибок
    setupErrorHandling();
}

/**
 * Настройка обработки ошибок
 */
function setupErrorHandling() {
    // Глобальный обработчик ошибок
    window.addEventListener('error', function (event) {
        console.error('Глобальная ошибка:', event.error);
        blogApi.showNotification('Произошла ошибка в приложении', 'error');
    });

    // Обработчик необработанных обещаний
    window.addEventListener('unhandledrejection', function (event) {
        console.error('Необработанное обещание:', event.reason);
        blogApi.showNotification('Произошла ошибка в приложении', 'error');
    });
}

/**
 * Показать информацию о приложении
 */
function showAppInfo() {
    const appInfo = {
        name: 'Блог на Go',
        version: '1.0.0',
        description: 'Простое веб-приложение для управления записями блога',
        author: 'Разработчик',
        backend: 'Go RESTful API',
        frontend: 'HTML, CSS, JavaScript, Bootstrap'
    };

    console.log('Информация о приложении:', appInfo);
    return appInfo;
}

// Инициализация приложения при загрузке
document.addEventListener('DOMContentLoaded', function () {
    showAppInfo();
});

// Экспорт основных функций
window.blogApp = {
    initApp,
    checkApiAvailability,
    setupErrorHandling,
    showAppInfo
};
