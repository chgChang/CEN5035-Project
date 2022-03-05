describe('login', function(){
    this.beforeEach(()=>{
        cy.visit('http://localhost:8000/user/login')
    })

    it('correctTest', function(){
        cy.get('.ant-pro-form-login-main').get('#email').type('11@111.com')
        cy.get('.ant-pro-form-login-main').get('#password').type('123456')
        cy.get('button:contains("Login")').click()
        cy.url().should('eq', 'http://localhost:8000/')//url equal to http://localhost:8000/
    })

    it('incorrectTest', function(){
        cy.get('.ant-pro-form-login-main').get('#email').type('11@111.com')
        cy.get('.ant-pro-form-login-main').get('#password').type('1234')
        cy.get('button:contains("Login")').click()
        cy.url().should('include', '/user/login')
        // cy.url().should('eq', 'http://localhost:8000/user/login')
    })

})