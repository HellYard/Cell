<?php
/**
 * Created by Daniel Vidmar.
 * User: Daniel
 * Date: 11/3/2016
 * Time: 12:56 AM
 * Version: Beta 2
 * Last Modified: 11/3/2016 at 12:56 AM
 * Last Modified by Daniel Vidmar.
 */

namespace lexer\token;


abstract class TokenType
{
  public abstract function getPattern();
}