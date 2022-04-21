describe('Checkout', function(){
    it('checkout', function(){

        //login
        cy.visit('http://localhost:8000/user/login')
        cy.get('.ant-pro-form-login-main').get('#email').type('cz@cz.com')
        cy.get('.ant-pro-form-login-main').get('#password').type('123456')
        cy.get('button:contains("Login")').click()

        //add elements to cart
        cy.get('#root > div > section > div > main > div > div.ant-pro-grid-content > div > div > div > div > div > div > div > div > div:nth-child(1) > div')
          .contains("Add to Cart").click()
        cy.get('#root > div > section > div > main > div > div.ant-pro-grid-content > div > div > div > div > div > div > div > div > div:nth-child(1) > div')
          .contains("Add to Cart").click()
        cy.get('#root > div > section > div.ant-layout > main > div > div.ant-pro-grid-content > div > div > div > div > div > div > div > div > div:nth-child(2)')
          .contains("Add to Cart").click()
        
        //go to checkout page
        cy.contains('Cart').click()
        cy.get('.ant-pro-grid-content').get('button:contains("checkout")').click()
        cy.wait(500)
        cy.contains('Next').click()
        
        cy.contains('Card number').type('123456')
        cy.contains('Name').type('zc')
        cy.contains('CVV').type('123')
        cy.contains('Next').click()

        cy.contains('name').type('Chang Zhou')
        cy.contains('Address').type('Gainesville')
        cy.contains('Code').type('12345')
        cy.contains('Next').click()
        cy.contains('Next').click()
        cy.contains('success')
    })

})