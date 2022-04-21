describe('register', function(){
    this.beforeEach(()=>{
        cy.visit('http://localhost:8000/user/login')
        cy.get('.ant-pro-form-login-main').get('#email').type('cz@cz.com')
        cy.get('.ant-pro-form-login-main').get('#password').type('123456')
        cy.get('button:contains("Login")').click()
    })

    it('button', function(){
        // cy.contains('Cart').should('have.attr', 'href', '/cart-list').should('have.attr', 'target', '_blank')
        cy.contains('Cart').then(link => {
            cy.request(link.prop('href')).its('status').should('eq', 200)
        })
        cy.contains('History').then(link => {
            cy.request(link.prop('href')).its('status').should('eq', 200)
        })
        cy.contains('Welcome').then(link => {
            cy.request(link.prop('href')).its('status').should('eq', 200)
        })

        cy.contains('Cart').click()
        cy.url().should('include', 'cart-list')

        cy.contains('History').click()
        cy.url().should('include', 'history')

        cy.contains("Welcome").click()
        cy.url().should('include', 'welcome')
    })

})