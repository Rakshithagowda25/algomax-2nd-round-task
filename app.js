// app.js
document.addEventListener('DOMContentLoaded', function () {
    // Fetch and display jobs
    axios.get('http://localhost:8080/jobs')
        .then(response => {
            const jobList = document.getElementById('job-list');
            response.data.forEach(job => {
                const listItem = document.createElement('li');
                listItem.className = 'list-group-item';
                listItem.textContent = `${job.title} at ${job.company} (${job.location})`;
                jobList.appendChild(listItem);
            });
        })
        .catch(error => console.error('Error fetching jobs:', error));

    // Handle job form submission
    const jobForm = document.getElementById('job-form');
    jobForm.addEventListener('submit', function (event) {
        event.preventDefault();

        const formData = new FormData(jobForm);
        const jobData = {};
        formData.forEach((value, key) => {
            jobData[key] = value;
        });

        axios.post('http://localhost:8080/jobs', jobData)
            .then(response => {
                // Refresh job list
                location.reload();
            })
            .catch(error => console.error('Error posting job:', error));
    });
});
