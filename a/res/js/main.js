var createExpenseButton = document.getElementById('add-expense-button');
createExpenseButton.addEventListener('click', function(){
	createExpenseButton.style.display = 'none';
	document.getElementById('add-expense-form').style.display = 'block';
});

getExpenses();

function getExpenses() {
	var timeline = document.getElementById('timeline-content');
	
	fetch("/expense", {
		credentials: 'include'
	})
		.then(response => response.json())
		.then(expenses => {
			var counter = 0;
			expenses.forEach(expense => {
				counter += 1;
				
				var timelineExpense = document.createElement('div');
				timelineExpense.classList.add('timeline-expense');
				var container = document.createElement('div');
				if (counter % 2 === 0){
					container.classList.add('content-right-container');
				}
				else {
					container.classList.add('content-left-container');
				}
				var content = document.createElement('div');
				if (counter % 2 === 0){
					content.classList.add('content-right');
				}
				else {
					content.classList.add('content-left');
				}
				var description = document.createElement('p');
				var amount = document.createElement('span');
				var meta = document.createElement('div');
				meta.classList.add('meta-date');
				var date = document.createElement('span');
				date.classList.add('date');
				var month = document.createElement('span');
				month.classList.add('month');
				
				description.innerHTML = expense.description.String;
				amount.innerHTML = expense.amount;
				
				content.appendChild(description);
				content.appendChild(amount);
				container.appendChild(content);
				timelineExpense.appendChild(container);
				meta.appendChild(date);
				meta.appendChild(month);
				timelineExpense.appendChild(meta);
				timeline.appendChild(timelineExpense);
				
				
			})
		})
}
