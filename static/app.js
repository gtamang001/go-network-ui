document.addEventListener("DOMContentLoaded", function() {
    const refreshInterval = 3000;

    function loadResults() {
        fetch('/results')
            .then(response => response.text())
            .then(html => {
                const table = new DOMParser().parseFromString(html, 'text/html');
                const newRows = table.querySelector('tbody').innerHTML;
                document.querySelector('#results-table tbody').innerHTML = newRows;
            });
    }

    // Auto-refresh results table
    setInterval(loadResults, refreshInterval);

    // Show spinner on form submit
    document.querySelectorAll("form").forEach(form => {
        form.addEventListener("submit", () => {
            document.querySelector("#spinner").classList.remove("d-none");
        });
    });

    // Dark mode toggle
    document.querySelector("#darkModeToggle").addEventListener("click", () => {
        document.body.classList.toggle("bg-dark");
        document.body.classList.toggle("text-light");
        document.querySelectorAll(".card").forEach(card => {
            card.classList.toggle("bg-secondary");
            card.classList.toggle("text-white");
        });
    });
});
