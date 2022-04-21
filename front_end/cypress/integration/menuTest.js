describe('register', function(){
    this.beforeEach(()=>{
        cy.visit('http://localhost:8000/user/login')
        cy.get('.ant-tabs-tab:contains("Register")').click()
    })
    it('newTest', function(){
        cy.get('.ant-pro-form-login-main').get('#email').type('123@111.com')
        cy.get('.ant-pro-form-login-main').get('#username').type('cc')
        cy.get('.ant-pro-form-login-main').get('#registerPassword').type('zc1234')
        cy.get('.ant-pro-form-login-main').get('#confirmPassword').type('zc1234')
        cy.get('button:contains("Login")').click()
        cy.url().should('eq', 'http://localhost:8000/')
    })

    it('existTest', function(){
        cy.get('.ant-pro-form-login-main').get('#email').type('11@111.com')
        cy.get('.ant-pro-form-login-main').get('#username').type('cc')
        cy.get('.ant-pro-form-login-main').get('#registerPassword').type('zc1234')
        cy.get('.ant-pro-form-login-main').get('#confirmPassword').type('zc1234')
        cy.get('button:contains("Login")').click()
        cy.url().should('include', '/user/login')
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